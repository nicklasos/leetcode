<?php

class Solution
{
    public function twoSum($nums, $target): array
    {
        for ($i = 0; $i < count($nums) - 1; $i++) {
            for ($j = $i + 1; $j < count($nums); $j++) {
                if ($nums[$i] + $nums[$j] === $target) {
                    return [$i, $j];
                }
            }
        }

        return [];
    }
}

$s = new Solution();

$nums = [2,7,11,15];
$target = 9;

$r = $s->twoSum($nums, $target);

var_dump($r);
