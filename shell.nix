{
  pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
      import (fetchTree nixpkgs.locked) {
        overlays = [
          (import "${fetchTree gomod2nix.locked}/overlay.nix")
        ];
      }
  ),
  mkGoEnv ? pkgs.mkGoEnv,
  gomod2nix ? pkgs.gomod2nix,
  pre-commit-hooks,
  ...
}: let
  goEnv = mkGoEnv {pwd = ./.;};

  pyPackages = ps: with ps; [pkgs.python311Packages.mkdocs-material pkgs.python311Packages.mkdocs pkgs.python311Packages.mkdocs-material-extensions];

  pre-commit-check = pre-commit-hooks.lib.${pkgs.system}.run {
    src = ./.;
    hooks = {
      gofmt.enable = true;
      golangci-lint = {
        enable = true;
        name = "golangci-lint";
        description = "Lint my golang code";
        files = "\.go$";
        entry = "${pkgs.golangci-lint}/bin/golangci-lint run --new-from-rev HEAD --fix";
        pass_filenames = false;
      };
      goimports = {
        enable = true;
        name = "goimports";
        description = "Format my golang code";
        files = "\.go$";
        entry = let
          script = pkgs.writeShellScript "precommit-goimports" ''
            set -e
            failed=false
            for file in "$@"; do
                # redirect stderr so that violations and summaries are properly interleaved.
                if ! ${pkgs.gotools}/bin/goimports -l -d "$file" 2>&1
                then
                    failed=true
                fi
            done
            if [[ $failed == "true" ]]; then
                exit 1
            fi
          '';
        in
          builtins.toString script;
      };

      commitizen = {
        enable = true;
        name = "commitizen";
        description = "Commit using conventional commits";
        entry = "${pkgs.commitizen}/bin/cz";
        require_serial = true;
        pass_filenames = false;
      };
    };
  };
in
  pkgs.mkShell {
    inherit (pre-commit-check) shellHook;
    packages = [
      goEnv
      gomod2nix
      pkgs.golangci-lint
      pkgs.go_1_23
      pkgs.gotools
      pkgs.go-junit-report
      pkgs.commitizen
      pkgs.just

      # vscode needed things
      pkgs.gopls

      # GIS/related to gen
      pkgs.qgis-ltr
      pkgs.imagemagick

      # Documentation
      pkgs.python311
      (pkgs.python311.withPackages pyPackages)
    ];

    buildInputs = [(pkgs.callPackage ./pkgs/worldpainter.nix {}) (pkgs.callPackage ./pkgs/minutor.nix {})];
    require_serial = true;
  }
