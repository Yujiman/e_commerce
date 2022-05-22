#######################  GO  ##########################
run:
	find  ./ -mindepth 1 -maxdepth 1  -not -path "./.*" -not -path "./resources" -not -path "./teamCity" -type d -print -exec make -s run -C {} \;
ci-cd:
	find  ./ -mindepth 1 -maxdepth 1  -not -path "./.*" -not -path "./resources" -not -path "./teamCity" -type d -print -exec make -s ci-cd -C {} \;
stop:
	for pid in $$(lsof ./*/ | grep "^main" | awk '{print $2}'); do kill -9 $${pid} || true; done

restart: stop run
