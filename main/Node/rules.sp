F * 0 => 0
0 * F => 0
1 * F => F
F * 1 => F
// Multiplications ^



// Divisions
C / 1 => 1
V / 1 => V
F / 0 => undef
F / inf => 0
inf / inf => undef


// Addition and Subtractions
0 + F => F
F + 0 => F
F - 0 => F
//0 - F => -F


// Undefined Exponents
0 ^ 0 => undef
F ^ 1 => F
F ^ 0 => 1
F ^ inf => inf
//F ^ -inf => 0


// Numeric Constant operations
N + N_2 => Addition()
N - N_2 => Subtraction()
N * N_2 => Multiplication()
N / N_2 => Division()
N ^ N_2 => Exponential()


// Commutative
V * C => C * V
(C * V) * C_1 => C * C_1 * V
C * (V * C_1) => C * C_1 * V


// Exponent rule
C * C => C ^ 2
V * V => V ^ 2
(V ^ C) * V => V ^ (C + 1)
V * (V ^ C) => V ^ (C + 1)
(C ^ N) * C => C ^ (N + 1)
C * (C ^ N) => C ^ (N + 1)

// Derivatives
D(C) => 0
D(V) => VarDerivative()
D(C * V) => C*D(V)
D(F * F_2) => D(F) * F_2 + F * D(F_2)
D(F + F_2) => D(F) + D(F_2)
D(F - F_2) => D(F) - D(F_2)
D(F ^ F_2) => F_2 * (F ^ F_2 - 1) D(F) + (F ^ F_2) * log(F) * D(F_2)
D(F / F_2) => (D(F) * F_2 - F * D(F_2))/(F_2^2)
D(sin(F)) => D(F) * cos(F)
D(cos(F)) => 0 - D(F) * sin(F)
D(log(F)) => D(F) / F
D(dV) => dV ^ 2


//i ^ 2 => -1
//i ^ C => imaginaryExponents()


// Trig Rules
sin(0) => 0
sin(pi) => 0
cos(0) => 1
//cos(pi) => -1

// Log rules
log(e) => 1
log(F ^ F_2) => F_2 * log(F)
log(F * F_2) => log(F) + log(F_2)
log(F / F_2) => log(F) - log(F_2)

