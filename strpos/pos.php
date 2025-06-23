<?php

class Solution {
    /**
     * @param String $haystack
     * @param String $needle
     * @return Integer
     */
    function strStr($haystack, $needle) {
        $res = strpos($haystack, $needle);

        if ($res === false) {
            return -1;
        }

        return $res;
    }
}

$s = new Solution();

var_dump($s->strStr("leetcode", "leeto"));