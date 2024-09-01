echo Expected result from: ""hello" standard | cat -e""
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $"
echo Results from tested program:
go run . "hello" standard | cat -e
echo "Press enter to continue"
read ans
echo Expected result from: ""hello world" shadow | cat -e""
                                                                                        $
_|                _| _|                                                     _|       _| $
_|_|_|     _|_|   _| _|   _|_|         _|      _|      _|   _|_|   _|  _|_| _|   _|_|_| $
_|    _| _|_|_|_| _| _| _|    _|       _|      _|      _| _|    _| _|_|     _| _|    _| $
_|    _| _|       _| _| _|    _|         _|  _|  _|  _|   _|    _| _|       _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|             _|      _|       _|_|   _|       _|   _|_|_| $
                                                                                        $
                                                                                        $"
echo
echo Results from tested program:
go run . "hello world" shadow | cat -e
echo "Press enter to continue"
read ans
echo Expected result from: ""nice 2 meet you" thinkertoy | cat -e""
                                                                       $
                       --                       o                      $
     o                o  o                      |                      $
o-o     o-o o-o         /        o-O-o o-o o-o -o-       o  o o-o o  o $
|  | | |    |-'        /         | | | |-' |-'  |        |  | | | |  | $
o  o |  o-o o-o       o--o       o o o o-o o-o  o        o--O o-o o--o $
                                                            |          $
                                                         o--o          $"
echo Results from tested program:
go run . "nice 2 meet you" thinkertoy | cat -e 
echo "Press enter to continue"
read ans
echo Expected result from:"\"#$%&/()*+,-./" "thinkertoy | cat -e""
o o         | |                                                  $
| |  | |   -O-O-      O          o  / \  o | o                 o $
    -O-O- o | |   o  /    o     /  o   o  \|/   |             /  $
     | |   -O-O-    /    /|    o   |   | --O-- -o-           o   $
    -O-O-   | | o  /  o o-O-  /    o   o  /|\   |    o-o    /    $
     | |   -O-O-  O       |  o      \ /  o | o     o     O o     $
            | |                                    |             $
                                                                 $"
echo Results from tested program:
go run . "\"#$%&/()*+,-./" thinkertoy | cat -e
