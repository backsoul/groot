git add .
git commit -m "Cambios y deploy"
git push origin master
ssh root@ssh.backsoul.dev "cd /var/docker-apps/groot && sudo sh update_branch.sh"