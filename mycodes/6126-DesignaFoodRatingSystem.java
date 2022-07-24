class FoodRatings {

    class Food {
        String name;
        String cuisine;
        int rating;

        public Food(String name, String cuisine, int rating) {
            this.name = name;
            this.cuisine = cuisine;
            this.rating = rating;
        }
    }

    Map<String, Food> map;
    Map<String, PriorityQueue<Food>> cuisineMap;

    public FoodRatings(String[] foods, String[] cuisines, int[] ratings) {
        map = new HashMap<>();
        cuisineMap = new HashMap<>();
        for (int i = 0; i < foods.length; i++) {
            Food food = new Food(foods[i], cuisines[i], ratings[i]);
            map.put(foods[i], food);
            if (!cuisineMap.containsKey(cuisines[i])) 
                cuisineMap.put(cuisines[i], new PriorityQueue<>(
                    (a, b) -> b.rating == a.rating ? a.name.compareTo(b.name) : b.rating - a.rating
                ));
            cuisineMap.get(cuisines[i]).offer(food);
        }
    }
    
    public void changeRating(String food, int newRating) {
        map.get(food).rating = newRating;
        PriorityQueue<Food> q = cuisineMap.get(map.get(food).cuisine);
        q.remove(map.get(food));
        q.offer(map.get(food));
    }
    
    public String highestRated(String cuisine) {
        return cuisineMap.get(cuisine).peek().name;
    }
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * FoodRatings obj = new FoodRatings(foods, cuisines, ratings);
 * obj.changeRating(food,newRating);
 * String param_2 = obj.highestRated(cuisine);
 */