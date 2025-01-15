import React, { useState, useEffect } from 'react';
import { View, Text, Modal, TouchableOpacity, FlatList, StyleSheet, Alert } from 'react-native';
import { Picker } from '@react-native-picker/picker';
import { Recipe, client, MealInput } from '../../Service/Api';
import SearchBar from '../Common/SearchBar';

interface AddMealFormProps {
  visible: boolean;
  onClose: () => void;
  selectedDate: string;
  onMealAdded: () => void;
}

const AddMealForm: React.FC<AddMealFormProps> = ({ visible, onClose, selectedDate, onMealAdded }) => {
  const [hour, setHour] = useState('12');
  const [minute, setMinute] = useState('00');
  const [guests, setGuests] = useState(1);
  const [recipeQuery, setRecipeQuery] = useState('');
  const [recipeSuggestions, setRecipeSuggestions] = useState<Recipe[]>([]);
  const [selectedRecipe, setSelectedRecipe] = useState<Recipe | null>(null);

  const fetchRecipeSuggestions = async (query: string) => {
    try {
      const data = await client.fetchRecipes({ name: query, preference: false });
      if (data.err) {
        throw new Error(`An error occurred while accessing the list of recipes : ${data.err}`);
      }   
      const filteredRecipes = data.recipes.filter(recipe =>
        recipe.name.toLowerCase().includes(query.toLowerCase())
      );
      setRecipeSuggestions(filteredRecipes);
    } catch (error) {
      console.error(error);
      if (error instanceof Error) {
        Alert.alert(error.message);
      } else {
        Alert.alert("Error", "An unexpected error occurred.");
      }
    }
  };

  useEffect(() => {
    if (recipeQuery.length > 1) {
      fetchRecipeSuggestions(recipeQuery);
    } else {
      setRecipeSuggestions([]);
    }
  }, [recipeQuery]);

  const handleSave = async () => {
    if (selectedRecipe) {
      const plannedAt: string = `${selectedDate}T${hour}:${minute}:00Z`;
      const mealData: MealInput = {
        planned_at: plannedAt,
        guests: guests,
        recipe_id: selectedRecipe.id
      };

      try {
        const response = await client.createMeal(mealData);
        if (response.err) {
          throw new Error(`An error occurred while accessing the list of recipes : ${response.err}`);
        }
        onMealAdded(); // Trigger parent callback
        onClose(); // Close this form
      } catch (error) {
        console.error(error);
        if (error instanceof Error) {
          Alert.alert(error.message);
        } else {
          Alert.alert("Error", "An unexpected error occurred.");
        }
      }
    } else {
      alert('Please select a recipe.');
    }
  };

  const handleRecipeSelect = (recipe: Recipe) => {
    setSelectedRecipe(recipe);
    setRecipeQuery(recipe.name);
    setRecipeSuggestions([]); // Clear suggestions after selection
  };

  return (
    <Modal visible={visible} transparent={true} animationType="slide">
      <View style={styles.modalContainer}>
        <View style={styles.modalContent}>
          <Text style={styles.modalTitle}>Ajout d'un repas</Text>

          <Text style={styles.label}>Heure</Text>
          <View style={styles.pickerContainer}>
            <Picker selectedValue={hour} style={styles.picker} onValueChange={(itemValue) => setHour(itemValue)}>
              {Array.from({ length: 24 }, (_, i) => (
                <Picker.Item key={i} label={`${i.toString().padStart(2, '0')} h`} value={i.toString().padStart(2, '0')} />
              ))}
            </Picker>
            <Picker selectedValue={minute} style={styles.picker} onValueChange={(itemValue) => setMinute(itemValue)}>
              {Array.from({ length: 12 }, (_, i) => (
                <Picker.Item key={i * 5} label={(i * 5).toString().padStart(2, '0')} value={(i * 5).toString().padStart(2, '0')} />
              ))}
            </Picker>
          </View>

          <Text style={styles.label}>Invités</Text>
          <Picker selectedValue={guests} style={styles.picker} onValueChange={(itemValue) => setGuests(itemValue)}>
            {Array.from({ length: 20 }, (_, i) => (
              <Picker.Item key={i + 1} label={`${i + 1}`} value={i + 1} />
            ))}
          </Picker>

          <Text style={styles.label}>Sélection de la recette</Text>
          <SearchBar
            value={recipeQuery}
            onChangeText={setRecipeQuery}
            placeholder="Search recipe..."
            style={styles.searchBar}
          />
          {recipeSuggestions.length > 0 && (
            <FlatList
              data={recipeSuggestions}
              keyExtractor={(item) => item.id.toString()}
              renderItem={({ item }) => (
                <TouchableOpacity style={styles.suggestionItem} onPress={() => handleRecipeSelect(item)}>
                  <Text>{item.name}</Text>
                </TouchableOpacity>
              )}
              style={styles.suggestionsList}
            />
          )}

          <TouchableOpacity onPress={handleSave} style={styles.saveButton}>
            <Text style={styles.saveButtonText}>Confirmer</Text>
          </TouchableOpacity>
          <TouchableOpacity onPress={onClose} style={styles.cancelButton}>
            <Text style={styles.cancelButtonText}>Annuler</Text>
          </TouchableOpacity>
        </View>
      </View>
    </Modal>
  );
};

const styles = StyleSheet.create({
  modalContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'rgba(0,0,0,0.5)',
  },
  modalContent: {
    width: '80%',
    backgroundColor: 'white',
    borderRadius: 10,
    padding: 20,
  },
  modalTitle: {
    fontSize: 20,
    fontWeight: 'bold',
    marginBottom: 10,
  },
  label: {
    fontSize: 16,
    marginVertical: 10,
  },
  pickerContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
  picker: {
    flex: 1,
  },
  searchBar: {
    marginBottom: 10,
  },
  suggestionItem: {
    padding: 10,
    backgroundColor: '#f0f0f0',
    borderBottomWidth: 1,
    borderBottomColor: '#ddd',
  },
  saveButton: {
    marginTop: 20,
    backgroundColor: 'green',
    paddingVertical: 10,
    borderRadius: 5,
    alignItems: 'center',
  },
  saveButtonText: {
    color: 'white',
    fontSize: 16,
  },
  cancelButton: {
    marginTop: 10,
    backgroundColor: 'red',
    paddingVertical: 10,
    borderRadius: 5,
    alignItems: 'center',
  },
  cancelButtonText: {
    color: 'white',
    fontSize: 16,
  },
  suggestionsList: {
    maxHeight: 150,
  },
});

export default AddMealForm;
