use strict;
use warnings;

my @nums = (2,7,11,15);
my $target = 9;

# Brute force - nested for loops are not a good idea
# for (my $i = 0; $i < scalar @nums - 1; $i++) {
#     for (my $j = $i + 1; $j < scalar @nums; $j++) {
#         if ( $nums[$i] + $nums[$j] == $target ){
#             print "[ $i, $j ]\n";
#             exit;
#         }
#     }
# }

# Using a hash
my %passedNums = ();

for (my $i = 0; $i < scalar @nums; $i++) {
    if ( defined $passedNums{ $target - $nums[$i] } ) {
        print "[ $i, $passedNums{ $target - $nums[$i] } ]\n";
    }
    $passedNums{ $nums[$i] } = $i;
}