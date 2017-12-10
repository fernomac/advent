(ns day3)

(def input 368078)

; expands a pair of directions and a count into a sequence of steps in the
; given directions.
(defn expand [[d1 d2] count]
  (concat (repeat count d1)
          (repeat count d2)))

; an infinite sequence of directions around the spiral
; (:right :up :left :left :down :down :right :right :right ...)
(def directions 
  (let [dirs (cycle [[:right :up] [:left :down]])
        counts (map inc (range))]
  (apply concat (map expand dirs counts))))

; given a coordinate and a direction, return the coordinate of the
; square in the given direction.
(defn nextcoord [coord dir]
  (let [key  (case dir (:left :right) :x (:up :down) :y)
        diff (case dir (:left :down) -1 (:right :up) 1)]
    (update coord key + diff)))

; returns the coordinates of the square with index n.
(defn coords [n]
  (reduce nextcoord {:x 0 :y 0} (take (dec n) directions)))

; the manhattan distance to the target square.
(def part-one
  (let [target (coords input)]
    (+ (Math/abs (:x target))
       (Math/abs (:y target)))))

(println part-one)
