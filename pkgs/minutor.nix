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
    version = "2.21.0";
    src = fetchFromGitHub {
      owner = "mrkite";
      repo = "minutor";
      rev = "2.21.0";
      sha256 = "0ldjnrk429ywf8cxdpjkam5k73s6fq7lvksandfn3xn7gl9np5rk";
    };

    nativeBuildInputs = [
      clang-tools
      vcpkg
      vcpkg-tool
      (
        # i work on my macbook and gdb is not available. script is linux anyway, so this is just for me :)
        if !isDarwin
        then gdb
        else null
      )
    ];

    buildInputs = [
      libsForQt5.qt5.qtbase
    ];

    configurePhase = ''
      qmake -makefile
    '';

    buildPhase = ''
      make -j$NIX_BUILD_CORES
    '';
    # fixes install phase error (no such file or directory)
    installPhase = ''
      mkdir -p $out/bin
      cp minutor $out/bin
    '';
  }
