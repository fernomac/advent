(def input (slurp "input.txt"))
(def lines (clojure.string/split-lines input))
(def matrix (map (fn [s] (map (fn [s] (Integer/parseInt s)) (clojure.string/split s #"\s+"))) lines))

(defn csum [row]
  (let [mx (apply max row) mn (apply min row)]
    (- mx mn)))

(def csums (map csum matrix))

(println csums)
(println (reduce + csums))
