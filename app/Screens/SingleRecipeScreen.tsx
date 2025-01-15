import React, { useEffect, useState } from "react";
import { View, Text, StyleSheet, ScrollView, TouchableOpacity } from "react-native";
import { useRoute } from "@react-navigation/native";
import { client, Recipe } from "../Service/Api";
import { SingleRecipeRouteProp } from "../Navigation/NavigationTypes";

export function SingleRecipeScreen() {
  const route = useRoute<SingleRecipeRouteProp>();
  const { id } = route.params;
  const [recipe, setRecipe] = useState<Recipe | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [guests, setGuests] = useState(1);

  useEffect(() => {
    const fetchRecipeDetails = async () => {
      try {
        const res = await client.fetchRecipeById(id, guests);
        setRecipe(res);
      } catch (err) {
        console.error("Error fetching recipe:", err);
        setError("Erreur lors de la récupération de la recette.");
      }
    };
    fetchRecipeDetails();
  }, [id]);

  const handleIncreaseGuests = () => {
    setGuests((prevGuests) => prevGuests + 1);
  };

  const handleDecreaseGuests = () => {
    setGuests((prevGuests) => (prevGuests > 1 ? prevGuests - 1 : 1));
  };

  if (!recipe) {
    return (
      <View style={styles.loadingContainer}>
        <Text style={styles.loadingText}>Loading...</Text>
      </View>
    );
  }

  const descriptionLines = recipe.description
    .split(".")
    .map((line) => line.trim())
    .filter((line) => line);

  return (
    <ScrollView style={styles.container}>
      <Text style={styles.name}>{recipe.name}</Text>
      <Text style={styles.details}>
        <Text style={styles.bold}>Temps de préparation:</Text>{" "}
        {recipe.preparation_time} min
      </Text>
      <Text style={styles.details}>
        <Text style={styles.bold}>Temps total:</Text> {recipe.total_time} min
      </Text>

      <View style={styles.guestsContainer}>
        <TouchableOpacity style={styles.button} onPress={handleDecreaseGuests}>
          <Text style={styles.buttonText}>-</Text>
        </TouchableOpacity>
        <Text style={styles.guestsText}>{guests} invité(s)</Text>
        <TouchableOpacity style={styles.button} onPress={handleIncreaseGuests}>
          <Text style={styles.buttonText}>+</Text>
        </TouchableOpacity>
      </View>
      <Text style={styles.details}>
        <Text style={styles.bold}>Ingrédients:</Text>
      </Text>
      {recipe.ingredients?.map((ingredient) => (
        <View key={ingredient.id} style={styles.ingredientContainer}>
          <Text style={styles.ingredientName}>{ingredient.name}</Text>
          <Text style={styles.ingredientQuantity}>
            {(ingredient.quantity * guests).toFixed(1)} {ingredient.unit}
          </Text>
        </View>
      ))}

      <Text style={styles.details}>
        <Text style={styles.bold}>Préparation:</Text>
      </Text>
      {descriptionLines.map((line, index) => (
        <Text key={index} style={styles.details}>
          - {line}.
        </Text>
      ))}
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 20,
    backgroundColor: "#FFF5E5",
  },
  loadingContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "#FFF5E5",
  },
  loadingText: {
    fontSize: 16,
    color: "#000",
  },
  name: {
    fontSize: 24,
    fontWeight: "bold",
    marginBottom: 10,
    color: "#000",
  },
  details: {
    fontSize: 16,
    marginBottom: 5,
    color: "#000",
  },
  bold: {
    fontWeight: "bold",
    color: "#000",
  },
  guestsContainer: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
    marginVertical: 10,
  },
  guestsText: {
    fontSize: 18,
    marginHorizontal: 10,
    color: "#000",
  },
  button: {
    backgroundColor: "#FFD700",
    padding: 10,
    borderRadius: 5,
  },
  buttonText: {
    fontSize: 18,
    color: "#000",
  },
  ingredientContainer: {
    flexDirection: "row",
    justifyContent: "space-between",
    marginVertical: 5,
  },
  ingredientName: {
    fontSize: 16,
    color: "#000",
  },
  ingredientQuantity: {
    fontSize: 16,
    color: "#000",
  },
  errorText: {
    color: "red",
    textAlign: "center",
    marginTop: 20,
  },
});