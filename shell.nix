{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/refs/tags/23.05.tar.gz") {} }:
pkgs.mkShell {
  nativeBuildInputs = with pkgs; [
    go gopls
  ];
}
