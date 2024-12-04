(defn executeMul [is] (* (read-string (nth is 1)) (read-string (nth is 2))))

(defn executeInstructions [is]
  (if (nil? (first is)) 0
      (case (first (first is))
        "do()" (executeInstructions (rest is))
        "don't()" (executeInstructions (drop-while #(not (= "do()" (first %))) is))
        (+ (executeMul (first is)) (executeInstructions (rest is))))))

(prn "Part 1" (reduce + (map executeMul (re-seq #"mul\((\d+),(\d+\))" (slurp "input.txt")))))
(prn "Part 2" (executeInstructions (re-seq #"mul\((\d+),(\d+\))|do\(\)|don\'t\(\)" (slurp "input.txt"))))
