(require '[clojure.string :as str])

(defn read-input [filename] (map #(map read-string (str/split % #"\s+")) (str/split-lines (slurp filename))))

(defn zip [sq] (map vector sq (rest sq)))
(defn diff [sq] (map #(reduce - (reverse %)) sq))

(defn incOrDec [d] (or (every? #(< 0 %) d) (every? #(> 0 %) d)))
(defn ruleMaxDiff [d] (every? #(and (<= (abs %) 3) (> (abs %) 0)) d))
(defn rules [d] (and (ruleMaxDiff d) (incOrDec d)))

(defn levelSequences [d] (map-indexed
                          (fn [i _] (concat (take i d) (drop (inc i) d)))
                          d))
(defn checkLevel [d] (rules (diff (zip d))))
(defn checkLevelSquences [d] (some true? (map checkLevel (levelSequences d))))

(defn runPart [f file] (count (filter true? (map f (read-input file)))))

(prn "Part1" (runPart checkLevel "input.txt"))
(prn "Part2" (runPart checkLevelSquences "input.txt"))

