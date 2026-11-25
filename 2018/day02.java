void main() {
    var counters = new HashMap<Integer, Integer>();
    var boxes = new BufferedReader(new InputStreamReader(System.in)).lines().toList();
    boxes.forEach(b -> {
        var letters = new HashMap<Integer, Integer>();
        b.chars().forEach(it -> letters.merge(it, 1, Integer::sum));
        letters.values().stream().distinct().forEach(c -> counters.merge(c, 1, Integer::sum));
    });
    IO.print(counters.get(2) * counters.get(3));
    for (int i = 1; i < boxes.size() - 1; i++) {
        for (int j = 0; j < i; j++) {
            var mismatches = new ArrayList<Integer>();
            var bi = boxes.get(i);
            var bj = boxes.get(j);
            for (int k = 0; k < Math.min(bi.length(), bj.length()) && mismatches.size() < 2; k++) {
                if (bi.charAt(k) != bj.charAt(k)) { mismatches.add(k); }
            }
            if (mismatches.size() == 1) {
                IO.println(" " + bi.substring(0, mismatches.get(0)) + bi.substring(mismatches.get(0) + 1));
                break;
            }
        }
    }
}