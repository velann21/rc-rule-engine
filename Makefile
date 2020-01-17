setup:
	if [ -e ./dist ]; then rm -rf ./dist; fi

dist:
	mkdir ./dist
	mkdir -p ./dist/{darwin/bin,linux/bin}

build: update
	GOOS=linux go build ./main.go;mv main ./dist/linux/bin/rc-rules-engine

docker-build-prep:
	cd docker; cp ../dist/linux/bin/rc-rules-engine .

docker: docker-build-prep
	cd docker; docker build -t rc-rules-engine:`cat .version` . --no-cache
	make docker-build-cleanup

docker-build-cleanup:
	rm docker/rc-rules-engine

release: update build docker-build-prep
	git tag `cat .version`
	git push --tags
	cd docker; docker build -t "staging.repo.rcplatform.io/reynencourt/rc-rules-engine:latest" .;
	cd docker; docker push "staging.repo.rcplatform.io/reynencourt/rc-rules-engine:latest"
	cd docker; docker tag "staging.repo.rcplatform.io/reynencourt/rc-rules-engine:latest" "staging.repo.rcplatform.io/reynencourt/rc-rules-engine:`cat ../.version`";
	cd docker; docker push "staging.repo.rcplatform.io/reynencourt/rc-rules-engine:`cat ../.version`"
	cd docker; rm -rf dist;
	make docker-build-cleanup

git-clean:
	git fetch -p && for branch in `git branch -vv | grep ': gone]' | awk '{print $1}'`; do git branch -D $branch; done

update:
	export GO111MODULE=on;go mod vendor

