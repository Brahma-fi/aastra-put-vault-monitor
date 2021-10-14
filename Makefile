image=brahma_aastra_put_monitor
tag=latest


build_d:
		docker build -t $(image) .

push: build_d
		docker tag $(image):$(tag) 820675130125.dkr.ecr.ap-south-1.amazonaws.com/$(image):latest
		aws ecr get-login-password --region ap-south-1 | docker login --username AWS --password-stdin 820675130125.dkr.ecr.ap-south-1.amazonaws.com
		docker push 820675130125.dkr.ecr.ap-south-1.amazonaws.com/$(image):$(tag)

run: 
		docker run -p 9000:8080 $(image):$(tag)

