echo  "Push any key to delete container or Ctrl+C to exit"
read  -n 1 -s key
docker rmi -f forum
ID=$(docker ps -a | awk '{print $1}' | head -n 2 | tail -n 1)
docker rm $ID
echo -e "\n"

echo "Image list:"
docker images
echo "-------------------------------------------------------"

echo "container list:"
docker ps -a 
echo  "-------------------------------------------------------------------------"