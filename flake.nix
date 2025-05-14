{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = { 
    self,
    nixpkgs,
    ...
  } @ inputs: let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
  in {

    devShells.${system}.default = pkgs.mkShell {
      
      nativeBuildInputs = [
        pkgs.go
        pkgs.gopls
        pkgs.golangci-lint
        pkgs.gnumake
      ];

      shellHook = ''
        echo "Welcome to the Go dev environment!"
        export GOPATH=$PWD/.gopath
        export GOBIN=$GOPATH/bin
        export PATH=$GOBIN:$PATH
        mkdir -p $GOBIN
      '';

    };

  };
}
