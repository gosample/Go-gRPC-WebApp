Author: Thejass Krishnan

Documentation
-------------

Completed Tasks:
- Rewrote the backend using gRPC and gRPC-gateway plugin
- Enabled TLS using self signed certificate
- Automated protobuf compilation
- Modified the UI using Bootstrap JS and Material Design
- Deployed on golang Docker container running on EC2 CoreOS instance 
- Automated the entire deployment process using Ansible

Deploying the source code:

	Use scripts/compile.sh for compiling the protobuf files.

	Create EC2 instance (CoreOS) and copy the files from local machine (Needs AWS access key, secret key & SSH key!) using this playbook. Your working directory should be stars-app/scripts/ while executing:

				sudo ansible-playbook playbook.yaml

	The deployment is completely automated with Ansible 2.0. Pls make sure your SSH key is configured and your AWS IAM user uses the same SSH key. You will have to specify your SSH key in the ansible playbook.

Keys Removed:
-------------
scripts/playbook.yaml
	-aws_access_key
	-aws_secret_key

utils/gh.go
	-GITHUB_API_KEY