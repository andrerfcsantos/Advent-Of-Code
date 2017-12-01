#!/usr/bin/perl

$/='';
chomp($input = <>);

@input  = map(int,(split '', $input));
@input_enum = map( [$_, $input[$_]], 0 .. $#input);
$step = (scalar @input) / 2;

$sum, $sum2;

for (@input_enum) {
    ($i,$number) = @$_;
    if ($number == $input[($i+1) % (scalar @input)]){
        $sum += $number;
    }
    if ($number == $input[($i+$step) % (scalar @input)]){
        $sum2 += $number;
    }

}

print "Part 1: $sum\n";
print "Part 2: $sum2\n";




