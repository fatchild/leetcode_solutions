my $example1 = "III";
my $example2 = "LVIII";
my $example3 = "MCMXCIV";
my $example4 = "IV";
my $example5 = "IX";
my $example6 = "XL";
my $example7 = "XC";
my $example8 = "CD";
my $example9 = "CM";

sub romanToInt{
    my $s = shift @_;

    my $t = 0;

    my %map = ( "I" => 1,"V" => 5,"X" => 10,"L" => 50,"C" => 100,"D" => 500,"M" => 1000 );

    for (my $i = 0; $i < length $s; $i++){

        my ($c, $n) = ($map{substr($s, $i, 1)}, $map{substr($s, $i + 1, 1)});
        # $c < $n ? $t -= $c : $t += $c; 
        if ($c < $n){
            $t -= $c
        } else {
            $t += $c
        }
    }

    return $t
}

print STDERR romanToInt($example1), " - 3\n";
print STDERR romanToInt($example2), " - 58\n";
print STDERR romanToInt($example3), " - 1994\n";
print STDERR romanToInt($example4), " - 4\n";
print STDERR romanToInt($example5), " - 9\n";
print STDERR romanToInt($example6), " - 40\n";
print STDERR romanToInt($example7), " - 90\n";
print STDERR romanToInt($example8), " - 400\n";
print STDERR romanToInt($example9), " - 900\n";