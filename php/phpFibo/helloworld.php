<?php
set_time_limit(0);
function Fibonacci($n)
{
    if ($n == 1 || $n == 2) {
        return 1;
    }else 
    return Fibonacci($n-1) + Fibonacci($n-2);
}

function FiboTime($fiboIndex){
	$startTime = microtime(true);
	$Fibo = Fibonacci($fiboIndex);
	$endTime = microtime(true);
	$totalTime = ($endTime - $startTime)*1000;
	return $totalTime;
}

$totalTime = 0;
$fiboIndex = 35;
echo 'The ' . $fiboIndex . 'th number in Fibonacci<br>';
for($flag = 1; $flag <= 8; $flag++)
{
	$totalTime = FiboTime($fiboIndex);
	echo $totalTime . '<br>';
	sleep(3);
}
echo 'success';


?>