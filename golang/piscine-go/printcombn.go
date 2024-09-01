package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	switch n {
	case 1:
		for a := '0'; a <= '9'; a++ {
			if a == '9' {
				z01.PrintRune(a)
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(a)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	case 2:
		for a := '0'; a <= '9'; a++ {
			for b := a + 1; b <= '9'; b++ {
				if a == '8' && b == '9' {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	case 3:
		for a := '0'; a <= '9'; a++ {
			for b := a + 1; b <= '9'; b++ {
				for c := b + 1; c <= '9'; c++ {
					if a == '7' && b == '8' && c == '9' {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						z01.PrintRune('\n')
					} else {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	case 4:
		for a := '0'; a <= '9'; a++ {
			for b := a + 1; b <= '9'; b++ {
				for c := b + 1; c <= '9'; c++ {
					for d := c + 1; d <= '9'; d++ {
						if a == '6' && b == '7' && c == '8' && d == '9' {
							z01.PrintRune(a)
							z01.PrintRune(b)
							z01.PrintRune(c)
							z01.PrintRune(d)
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(a)
							z01.PrintRune(b)
							z01.PrintRune(c)
							z01.PrintRune(d)
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	case 5:
		for a := '0'; a <= '9'; a++ {
			for b := a + 1; b <= '9'; b++ {
				for c := b + 1; c <= '9'; c++ {
					for d := c + 1; d <= '9'; d++ {
						for e := d + 1; e <= '9'; e++ {
							if a == '5' && b == '6' && c == '7' && d == '8' && e == '9' {
								z01.PrintRune(a)
								z01.PrintRune(b)
								z01.PrintRune(c)
								z01.PrintRune(d)
								z01.PrintRune(e)
								z01.PrintRune('\n')
							} else {
								z01.PrintRune(a)
								z01.PrintRune(b)
								z01.PrintRune(c)
								z01.PrintRune(d)
								z01.PrintRune(e)
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	case 6:
		for a := '0'; a <= '9'; a++ {
			for b := a + 1; b <= '9'; b++ {
				for c := b + 1; c <= '9'; c++ {
					for d := c + 1; d <= '9'; d++ {
						for e := d + 1; e <= '9'; e++ {
							for f := e + 1; f <= '9'; f++ {
								if a == '4' && b == '5' && c == '6' && d == '7' && e == '8' && f == '9' {
									z01.PrintRune(a)
									z01.PrintRune(b)
									z01.PrintRune(c)
									z01.PrintRune(d)
									z01.PrintRune(e)
									z01.PrintRune(f)
									z01.PrintRune('\n')
								} else {
									z01.PrintRune(a)
									z01.PrintRune(b)
									z01.PrintRune(c)
									z01.PrintRune(d)
									z01.PrintRune(e)
									z01.PrintRune(f)
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
	case 8:
		for a := '0'; a <= '9'; a++ {
			for b := a + 1; b <= '9'; b++ {
				for c := b + 1; c <= '9'; c++ {
					for d := c + 1; d <= '9'; d++ {
						for e := d + 1; e <= '9'; e++ {
							for f := e + 1; f <= '9'; f++ {
								for g := f + 1; g <= '9'; g++ {
									for h := g + 1; h <= '9'; h++ {
										if a == '2' && b == '3' && c == '4' && d == '5' && e == '6' && f == '7' && g == '8' && h == '9' {
											z01.PrintRune(a)
											z01.PrintRune(b)
											z01.PrintRune(c)
											z01.PrintRune(d)
											z01.PrintRune(e)
											z01.PrintRune(f)
											z01.PrintRune(g)
											z01.PrintRune(h)
											z01.PrintRune('\n')
										} else {
											z01.PrintRune(a)
											z01.PrintRune(b)
											z01.PrintRune(c)
											z01.PrintRune(d)
											z01.PrintRune(e)
											z01.PrintRune(f)
											z01.PrintRune(g)
											z01.PrintRune(h)
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
	case 7:
		for a := '0'; a <= '9'; a++ {
			for b := a + 1; b <= '9'; b++ {
				for c := b + 1; c <= '9'; c++ {
					for d := c + 1; d <= '9'; d++ {
						for e := d + 1; e <= '9'; e++ {
							for f := e + 1; f <= '9'; f++ {
								for g := f + 1; g <= '9'; g++ {
									if a == '3' && b == '4' && c == '5' && d == '6' && e == '7' && f == '8' && g == '9' {
										z01.PrintRune(a)
										z01.PrintRune(b)
										z01.PrintRune(c)
										z01.PrintRune(d)
										z01.PrintRune(e)
										z01.PrintRune(f)
										z01.PrintRune(g)
										z01.PrintRune('\n')
									} else {
										z01.PrintRune(a)
										z01.PrintRune(b)
										z01.PrintRune(c)
										z01.PrintRune(d)
										z01.PrintRune(e)
										z01.PrintRune(f)
										z01.PrintRune(g)
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
	case 9:
		for a := '0'; a <= '9'; a++ {
			for b := a + 1; b <= '9'; b++ {
				for c := b + 1; c <= '9'; c++ {
					for d := c + 1; d <= '9'; d++ {
						for e := d + 1; e <= '9'; e++ {
							for f := e + 1; f <= '9'; f++ {
								for g := f + 1; g <= '9'; g++ {
									for h := g + 1; h <= '9'; h++ {
										for i := h + 1; i <= '9'; i++ {
											if a == '1' && b == '2' && c == '3' && d == '4' && e == '5' && f == '6' && g == '7' && h == '8' && i == '9' {
												z01.PrintRune(a)
												z01.PrintRune(b)
												z01.PrintRune(c)
												z01.PrintRune(d)
												z01.PrintRune(e)
												z01.PrintRune(f)
												z01.PrintRune(g)
												z01.PrintRune(h)
												z01.PrintRune(i)
												z01.PrintRune('\n')
											} else {
												z01.PrintRune(a)
												z01.PrintRune(b)
												z01.PrintRune(c)
												z01.PrintRune(d)
												z01.PrintRune(e)
												z01.PrintRune(f)
												z01.PrintRune(g)
												z01.PrintRune(h)
												z01.PrintRune(i)
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
