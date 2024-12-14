{
  description = "Flake for Private Alpha";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.11";
  inputs.nixpkgs-unstable.url = "github:NixOS/nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = {
    self,
    nixpkgs,
    nixpkgs-unstable,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem
    (
      system: let
        stable = nixpkgs.legacyPackages.${system};
        unstable = nixpkgs-unstable.legacyPackages.${system};
      in {
        devShells.default = stable.mkShell {
          packages = with stable; [
            go
            golangci-lint
            gomodifytags
            gopls
            gotools
            go-tools
            gotests
            delve
            impl
            gotestsum
          ];
        };
      }
    );
}
