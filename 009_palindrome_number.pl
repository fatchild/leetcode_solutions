use POSIX;

sub isPalindrome{
    my $number = shift @_;

    return 0 if $number < 0 or ($number % 10 == 0 && $number != 0);

    my $reverse = 0;

    for (my $i = $number; $i > 0; $i = floor($i / 10)){
        my $lastDigit = $i % 10; # Break off the last digit

        $reverse = ($reverse * 10) + $lastDigit; # Append the last digit to the reverse
    }

    $number == $reverse ? return 1 : return 0;
}

print STDERR isPalindrome(121), "\n";
print STDERR isPalindrome(-121), "\n";
print STDERR isPalindrome(10), "\n";

