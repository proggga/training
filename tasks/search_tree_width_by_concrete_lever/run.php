<?php


class Node {
  private $childrens = [];

  public $value = null;

  public function addChild(Node $node) {
    $this->childrens[] = $node;
  }

  public function getChildrens() {
    return $this->childrens;
  }

  public function __construct($val) {
    $this->value = $val;
  }
}

function recursionWidthSearch(Node $node, $required_level, $level = 0) {
  if ($required_level < 1) {
    return 1;
  }
  if ($level > ($required_level - 1)) {
    return 0;
  }
  $sum = 0;
  foreach ($node->getChildrens() as $child) {
    $sum += recursionWidthSearch($child, $required_level, $level + 1);
  }
  if ($level === ($required_level - 1)) {
    $sum += count($node->getChildrens());
  }
  return $sum;
}


function main() {
  $root = Helper::genTree(6, 4, 12);
  Helper::printTree($root);
  $found_zero = false;
  $level          = 0;
  $expected_array = [];
  while (!$found_zero) {
    $res = recursionWidthSearch($root, $level);
    $expected_array[] = $res;
    $found_zero = $res === 0;
    $level++;
  }
  if ($expected_array !== [1, 3, 4, 4, 5, 3, 1, 0]) {
    echo "Code failed\n";
  }else{
    echo "Code success\n";
  }
}

main();


class Helper {
  static $last_id = 0;
  private static function getNode() {
    return new Node(++self::$last_id);
  }

  private static function getRoot(Node $parent, $levels, $max_childs, $level = 0) {
    if ($level >= $levels) {
      return $parent;
    }
    $full_chance = 1 / 3;
    $childes     = max(mt_rand(1, ceil(1 / $full_chance)) === 1 ? $max_childs : $max_childs - $level, 0);
    $nodes       = mt_rand(0, $childes);
    for ($i = 0; $i < $nodes; $i++) {
      $node = self::getNode();
      self::getRoot($node, $levels, $max_childs, $level + 1);
      $parent->addChild($node);
    }
    return $parent;
  }

  public static function genTree($levels, $max_childs, $srand_seed) {
    srand($srand_seed);
    return self::getRoot(self::getNode(), $levels, $max_childs);
  }

  public static function printTree(Node $node, $extra_offset = 0, $level = 0) {
    $len = 3;
    echo ($level >= 1 ? str_repeat('|  ' . str_repeat(' ', $len), $level - 1) . '|-' : '') . sprintf("%{$len}d", $node->value);

    if ($node->getChildrens()) {
      echo "-+\n";
      foreach ($node->getChildrens() as $children) {
        self::printTree($children, $extra_offset, $level + 1);
      }
    } else {
      echo "\n";
    }
  }
}