{
  description = "DevOps Go CLI";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        packages = rec {
          default = pkgs.buildGoModule {
            pname = "cli";
            version = "1.1.1";
            src = ./.;
            vendorHash = "sha256-Y/OQOYzCZ6Rv1T3FDWfFhdgskd1Ugv+RIfL6GCkDaOg=";
          };
        };
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs;[
            go
            golangci-lint
            golangci-lint-langserver
            pre-commit
            gnumake
            git
            curl
            markdownlint-cli
            gitleaks
            yamllint
            shellcheck
          ];
          shellHook = ''
            echo "Running CLI development environment"
            echo "Go: $(go version)"

            if [ -d .git ] && [ ! -f .git/hooks/pre-commit ]; then
              echo "Install pre-commit hooks"
              pre-commit install
            fi
            cleanup() {
              rm -rf result
              # rm -rf dist/ build/
            }
            trap cleanup EXIT
          '';
        };
      });
}
