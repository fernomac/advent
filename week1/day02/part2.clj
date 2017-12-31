(ns day2
 (:require [clojure.math.combinatorics :as combo]))

(def input (slurp "input.txt"))
(def lines (clojure.string/split-lines input))
(def matrix (map (fn [s] (map (fn [s] (Integer/parseInt s)) (clojure.string/split s #"\s+"))) lines))

(def combos (map (fn [row] (combo/combinations row 2)) matrix))

(defn divides [a b]
  (if (> a b)
    (== (mod a b) 0)
    (== (mod b a) 0)))

(defn divisor [a b]
  (if (> a b)
    (/ a b)
    (/ b a)))

(defn divisor-or-zero [pair]
  (let [a (nth pair 0) b (nth pair 1)]
    (if (divides a b)
      (divisor a b)
      0)))

(def divisors
  (map (fn [row]
    (map divisor-or-zero row))
    combos))

(def reduced (map (fn [row] (reduce + row)) divisors))

(println divisors)
(println reduced)
(println (reduce + reduced))
