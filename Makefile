shell:
	nix-shell
gen: shell
	go generate ./...