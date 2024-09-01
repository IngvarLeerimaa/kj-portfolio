echo Expected result from: "--output=test01.txt First\nTest shadow""
                                       $
_|_|_|_| _|                     _|     $
_|          _|  _|_|   _|_|_| _|_|_|_| $
_|_|_|   _| _|_|     _|_|       _|     $
_|       _| _|           _|_|   _|     $
_|       _| _|       _|_|_|       _|_| $
                                       $
                                       $
                                      $
_|_|_|_|_|                     _|     $
    _|       _|_|     _|_|_| _|_|_|_| $
    _|     _|_|_|_| _|_|       _|     $
    _|     _|           _|_|   _|     $
    _|       _|_|_| _|_|_|       _|_| $
                                      $
                                      $
"
echo Results from tested program:
go run . --output=test01.txt "First\nTest" shadow
cat test01.txt -e
echo "Press enter to continue"
read ans
echo Expected result from: --output=test02.txt "123 -> \"#$%@"" thinkertoy
                                    o o         | |               $
  0    --  o-o            o         | |  | |   -O-O-      O   o   $
 /|   o  o    |            \            -O-O- o | |   o  /   / \  $
o |     /   oo              O            | |   -O-O-    /   o O-o $
  |    /      |       o-o  /            -O-O-   | | o  /  o  \    $
o-o-o o--o o-o            o              | |   -O-O-  O       o-  $
                                                | |               $
                                                                  $
"
echo Results from tested program:
go run . --output=test02.txt "123 -> \"#$%@" thinkertoy
cat test02.txt -e
echo "Press enter to continue"
read ans
echo Expected result from: --output=test03.txt "432 -> #$%&@" shadow"
                                                                                                                  $
_|  _|   _|_|_|     _|_|                    _|             _|  _|     _|   _|_|    _|   _|           _|_|_|_|_|   $
_|  _|         _| _|    _|                    _|         _|_|_|_|_| _|_|_| _|_|  _|   _|  _|       _|          _| $
_|_|_|_|   _|_|       _|         _|_|_|_|_|     _|         _|  _|   _|_|       _|       _|_|  _| _|    _|_|_|  _| $
    _|         _|   _|                        _|         _|_|_|_|_|   _|_|   _|  _|_| _|    _|   _|  _|    _|  _| $
    _|   _|_|_|   _|_|_|_|                  _|             _|  _|   _|_|_| _|    _|_|   _|_|  _| _|    _|_|_|_|   $
                                                                      _|                           _|             $
                                                                                                     _|_|_|_|_|_| $
"
echo Results from tested program:
go run . --output=test03.txt "432 -> #$%&@" shadow
cat test03.txt -e
echo "Press enter to continue"
read ans
echo Expected result from:go run . --output=test04.txt "hello""
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $"
echo Results from tested program:
go run . --output=test04.txt "hello"
cat test04.txt -e
