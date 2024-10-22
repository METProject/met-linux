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
  stdenv,
  ...
}: let
  goEnv = mkGoEnv {pwd = ./.;};

  # Python 3.13
  # Nixpkgs does not enable LTO nor optimizations for reproducability, however we need it for some of the scripts to not be that slow...
  # However, building will be slow. oh well
  python313Optimized = pkgs.python313.override {
    enableOptimizations = true;
    enableLTO = true;
    reproducibleBuild = false;
    enableGIL = false;
    self = python313Optimized;
  };

  pyPackages = ps: with ps; [
    pkgs.python313Packages.mkdocs-material 
    pkgs.python313Packages.mkdocs 
    pkgs.python313Packages.mkdocs-material-extensions
    
    
    ];

  pre-commit-check = pre-commit-hooks.lib.${pkgs.system}.run {
    src = ./.;
    hooks = {
      gofmt.enable = true;
      golangci-lint = {
        enable = false;
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
    packages = with pkgs; [
      goEnv
      gomod2nix
      golangci-lint
      go_1_23
      gotools
      go-junit-report
      commitizen
      just

      # vscode needed things
      gopls

      # GIS/related to gen
      # disabling qgis-ltr on macos; only for dev on my laptop
      (
        if !stdenv.isDarwin
        then qgis-ltr
        else null
      )
      imagemagick

      # Python
      python313Optimized
      (python313Optimized.withPackages pyPackages)
    ];

    buildInputs = [(pkgs.callPackage ./pkgs/worldpainter.nix {}) ( if !stdenv.isDarwin then (pkgs.callPackage ./pkgs/minutor.nix {}) else null)];
    require_serial = true;
  }
