{
  stdenv,
  jdk17_headless,
}:
stdenv.mkDerivation {
  pname = "worldpainter";
  version = "2.22.1";
  src = builtins.fetchurl {
    url = "https://www.worldpainter.net/files/worldpainter_2.22.1.tar.gz";
    sha256 = "0bh3y73ibzgyqg3aqmkx93b2vym2dga0ybf0r8a2q0rkf0wwrwcy";
  };

  buildInputs = [jdk17_headless];

  unpackPhase = ''
    # unpack the package
    tar -xzvf $src
  '';

  installPhase = ''
    # install the package
    mkdir -p $out/bin
    cp -r * $out/bin

    # modify vmoptions
    sed -i 's/# -Xmx512m/-Xmx6G/g' $out/bin/worldpainter/wpscript.vmoptions
  '';

  # add an alias to run the binary
  shellHook = ''
    export PATH=$out/bin:$PATH
  '';
}
