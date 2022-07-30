git add .
git commit -m "Cambios y deploy"
git push origin master
ssh backsoul@ssh.backsoul.xyz "cd /var/docker-apps/groot && sudo sh update_branch.sh"