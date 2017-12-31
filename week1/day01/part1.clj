(def input (slurp "input.txt"))
(def values 
 (vec
  (map
    (fn [c] (- (int c) (int \0)))
    (seq input)
  )
 )
)
(def shifted
 (conj
   (subvec values 1)
   (nth values 0)
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
