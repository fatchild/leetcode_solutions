use warnings;

sub longestCommonPrefix{
    my @strs = @_;

    return "" if ( scalar @strs == 0 || $strs[0] eq "" );
    return $strs[0] if (scalar @strs == 1);

    my $lcp = "";

    for (my $i = 0; $i < length $strs[0]; $i++){
        my $hold = substr($strs[0], $i, 1);

        for (my $j = 0; $j < scalar @strs; $j++){
            return $lcp if (!defined substr($strs[$j], $i, 1) or substr($strs[$j], $i, 1) ne $hold);
        }
        $lcp = $lcp.$hold;
    }
    return $lcp;
}


my @testString1 = ("flower","flow","flight");
my @testStrings2 = ("dog","racecar","car");


longestCommonPrefix(@testString1) eq "fl" ? print "PASSED\n" : print "FAILED\n";
longestCommonPrefix(@testString2) eq "" ? print "PASSED\n" : print "FAILED\n";
longestCommonPrefix(("")) eq "" ? print "PASSED\n" : print "FAILED\n";
longestCommonPrefix(("a")) eq "a" ? print "PASSED\n" : print "FAILED\n";
longestCommonPrefix((("flower","flower","flower","flower"))) eq "flower" ? print "PASSED\n" : print "FAILED\n";