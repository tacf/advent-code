(require '[clojure.string :as str])

(defn process-file [filename]
  (with-open [rdr (clojure.java.io/reader filename)]
    (let [rows  (map #(map read-string
                           (str/split % #"\s+"))
                       (line-seq rdr))]
      {:col1 (sort (map first rows))
       :col2 (sort (map second rows))})))


(let [{:keys [col1 col2]} (process-file "input.txt")]
  ; (println "Column 1:" col1)
  ; (println "Column 2:" col2)
  (println "Part1:" (reduce + (map #(Math/abs (- (second %) (first %))) (map vector col1 col2))))
  (println "Part2:" (reduce + (map (fn [e] (* e (count (filter #{e} col2)))) col1))))
