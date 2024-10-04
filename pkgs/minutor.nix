# TODO: Make this actually compile. Currently, its complaining because Minutor is usually a GUI app with CLI capabilities. They are the same executable. Also, this doesn't work on my macbook (expected).
# This project is not planned on supporting macOS, so I'm not going to spend time on it. Just annoying with my dev environment. Anyone who knows Nix and on a desktop, please help me out. :yikes:
{
  stdenv,
  fetchFromGitHub,
  clang-tools,
  vcpkg,
  vcpkg-tool,
  libsForQt5,
  gdb,
  ...
}: let
  isDarwin = stdenv.isDarwin;
in
  stdenv.mkDerivation {
    dontWrapQtApps = true;
    pname = "minutor";
    version = "2.20.0";
    src = fetchFromGitHub {
      owner = "mrkite";
      repo = "minutor";
      rev = "2.20.0";
      sha256 = "1i0g0vb1q4160mz787sv05j30531ih66b1m8hbb44srrzmihm1bi";
    };

    nativeBuildInputs = [
      clang-tools
      vcpkg
      vcpkg-tool
      (
        if !isDarwin
        then gdb
        else null
      )
    ];

    buildInputs = [
      libsForQt5.qt5.qtbase
    ];

    configurePhase = ''
      qmake
    '';

    buildPhase = ''
      make
    '';
  }
