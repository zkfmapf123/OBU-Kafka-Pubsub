up:
	@docker-compose up --build -d

scp-up:
	@instanceIP=$$(aws ec2 describe-instances \
	--region ap-northeast-2 \
	--filters "Name=tag:Name,Values=pubsub-test" \
	--query "Reservations[*].Instances[*].PublicIpAddress" \
	--output text); \
	scp -r ~/.ssh/id_rsa ./publish ./subscriber *.yml Makefile ec2-user@$$instanceIP:~/
