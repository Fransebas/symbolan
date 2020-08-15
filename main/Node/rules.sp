F * 0 => 0
0 * F => 0
1 * F => F
F * 1 => F


-(-F) => F
-1 * F => -F
F * -1 => -F
(-F_1) * (-F_2) => (F_1) * (F_2)
(-1) * (F) => -F
//// Multiplications ^


//// Divisions
C / 1 => 1
V / 1 => V
(V^C) / V => V^(C-1)
F / 0 => undef
F / inf => 0
inf / inf => undef


//// Addition and Subtractions
0 + F => F
F + 0 => F
F - 0 => F
0 - F => -F
C - C => 0
V - V => 0
N - N => 0
F + (-F_2) => F - F_2
(-F) + F_2 => F_2 - F

F - (-F_2) => F + F_2
(-F) - F_2 => -(F_2 + F)

//// Undefined Exponents
0 ^ 0 => undef
F ^ 1 => F
F ^ 0 => 1
F ^ inf => inf
F ^ -inf => 0


//// Numeric Constant operations
N * (-N_2) => -(N * N_2)
(-N) * N_2 => -(N * N_2)

N / (-N_2) => -(N / N_2)
(-N) / N_2 => -(N / N_2)


N + N_2 => Addition()
N - N_2 => Subtraction()
N * N_2 => Multiplication()
N / N_2 => Division()
N ^ N_2 => Exponential()
(-N) ^ N_2 => SignedExponential()
N ^ (-N_2) => SignedExponential()
(-N) ^ (-N_2) => SignedExponential()


//// Commutative Rule
//F * C => C * F
//(C * F) * C_1 => (C * C_1) * F
//C_1 * (C * F) => (C * C_1) * F
//(C * F)*(C_1 * F_1) => (C * C_1)*(F * F_1)

//F * N => N * F
(N * F) * N_1 => (N * N_1) * F
(N * F)*(N_1 * F_1) => (N * N_1)*(F * F_1)


// Exponent rule
C * C => C ^ 2
V * V => V ^ 2
(V ^ C) * V => V ^ (C + 1)
V * (V ^ C) => V ^ (C + 1)
(C ^ N) * C => C ^ (N + 1)
C * (C ^ N) => C ^ (N + 1)


i ^ 2 => -1
////i ^ C => imaginaryExponents()


// Trig Rules
sin(0) => 0
sin(pi) => 0
sin(pi/2) => 1
sin(-(pi/2)) => -1
cos(0) => 1
cos(pi) => -1
cos(pi/2) => 0
cos(-(pi/2)) => 0
(sin(F)^2) + (cos(F)^2) => 1
(cos(F)^2) + (sin(F)^2) => 1

// Log rules
log(e) => 1
log(F ^ F_2) => F_2 * log(F)
log(F * F_2) => log(F) + log(F_2)
log(F / F_2) => log(F) - log(F_2)



// Derivatives
D(-F) => -D(F)
D(C) => 0
D(V) => VarDerivative()
D(C * V) => C*D(V)
D(F * F_2) => D(F) * F_2 + F * D(F_2)
D(F + F_2) => D(F) + D(F_2)
D(F - F_2) => D(F) - D(F_2)
D(F ^ F_2) => F_2 * (F ^ (F_2 - 1)) * D(F) + (F ^ F_2) * log(F) * D(F_2)
D(F / F_2) => (D(F)/F_2) - ((F/(F_2^2))*D(F_2))

D(sin(F)) => cos(F) * D(F)
D(cos(F)) => sin(F) * D(F)
D(tan(F)) => ((1/cos(F))^2)*D(F)

D(sinh(F)) => cosh(F) * D(F)
D(cosh(F)) => sinh(F) * D(F)
D(tanh(F)) => (1 - tanh(F)^2) * D(F)

D(log(F)) => D(F) / F
D(dV) => dV ^ 2


////// Sorting Multiplication

!FN * FN => FN * !FN

(FN * !FN) * (FN_1 * !FN_3) => (FN * FN_1) * (!FN * !FN_3)

( !FN ) * (FN * !FN_3) => FN * (!FN * !FN_3)
(FN * !FN) * !FN_3 => FN * (!FN * !FN_3)
( FN ) * (FN_1 * !FN) => (FN * FN_1) * (!FN)

// Sorting Addition

!FN + FN => FN + !FN

(FN + !FN) + (FN_1 + !FN_3) => (FN + FN_1) + (!FN + !FN_3)

( !FN ) + (FN + !FN_3) => FN + (!FN + !FN_3)
(FN + !FN) + !FN_3 => FN + (!FN + !FN_3)
( FN ) + (FN_1 + !FN) => (FN + FN_1) + (!FN)