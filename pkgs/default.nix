let
  nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-24.05";
  pkgs = import nixpkgs {
    config = {};
    overlays = [];
  };
in {
  worldpainter = pkgs.callPackage ./worldpainter.nix {};
  minutor = pkgs.callPackage ./minutor.nix {};
}
