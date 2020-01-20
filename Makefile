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

push-rules:
	tar czf events_rules.tar.gz dsl
	ssh -i ~/rc-dev.pem ubuntu@ec2-18-216-29-160.us-east-2.compute.amazonaws.com "sudo rm -rf /tmp/*"
	scp -r -i ~/rc-dev.pem events_rules.tar.gz ubuntu@ec2-18-216-29-160.us-east-2.compute.amazonaws.com:/tmp/
	ssh -i ~/rc-dev.pem ubuntu@ec2-18-216-29-160.us-east-2.compute.amazonaws.com "sudo cp /tmp/events_rules.tar.gz /var/www/html/promrules/events_rules.tar.gz"

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

