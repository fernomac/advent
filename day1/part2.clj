(def input (slurp "input.txt"))
(def values 
 (vec
  (map
    (fn [c] (- (int c) (int \0)))
    (seq input)
  )
 )
)
(def halfway (/ (count values) 2))
(def shifted
 (concat
   (subvec values halfway)
   (subvec values 0 halfway)
 )
)
(def pairs (map vector values shifted))
(def result
 (reduce +
  (map
   (fn [v] (if (== (nth v 0) (nth v 1)) (nth v 0) 0))
   pairs
  )
 )
)
(println result)
