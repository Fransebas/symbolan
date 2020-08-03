F * 0 => 0
0 * F => 0

1 * F => F
F * 1 => F

0 + F => F
F + 0 => F

F / 0 => undef
F / inf => 0
inf / inf => undef

0 ^ 0 => undef
F ^ 1 => F
F ^ 0 => 1
F ^ inf => inf
F ^ -inf => 0


N + N_2 => addition()
N - N_2 => subtraction()
N * N_2 => multiplication()
N / N_2 => division()
N ^ N_2 => exponential()


V * V => V ^ 2
V * C => C * V

//i ^ 2 => -1
//i ^ C => imaginaryExponents()

sin(0) => 0
sin(pi) => 0
cos(0) => 1
//cos(pi) => -1
