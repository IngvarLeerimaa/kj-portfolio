echo -e "Run.sh will build the docker image.\n"
echo -e "If you want to delete the container and image, run removeDocker.sh\n"
echo  -e "After the build is complete push "Ctrl+C" to stop docker \n"
echo  -e "Docker port is mapped at http://localhost:5000"
echo  -e "------------------------------------------------------------------\n"
echo "Push any key to start building the docker image or ctrl+c to exit"
read  -n 1 -s
#building the docker image
docker build -t forum .

docker run --name forum-container -p 5000:8080 forum

echo -e "-------------------------------------------------------------\n"
