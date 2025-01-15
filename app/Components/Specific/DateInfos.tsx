import React, { useState } from 'react';
import { View, Text, Modal, TouchableOpacity, FlatList, StyleSheet } from 'react-native';
import AddMealForm from './AddMealForm';
import { Meal } from '../../Service/Api';
import { useNavigation } from "@react-navigation/native";
import { RecipesStackNavigationProp } from '../../Navigation/NavigationTypes';

interface DateInfosProps {
  visible: boolean;
  onClose: () => void;
  selectedDate: string;
  meals: Meal[];
}

const DateInfos: React.FC<DateInfosProps> = ({ visible, onClose, selectedDate, meals }) => {
  const [addMealVisible, setAddMealVisible] = useState(false);
  const navigation = useNavigation<RecipesStackNavigationProp>();

  const renderMealItem = ({ item }: { item: Meal }) => {
    const mealTime = new Date(item.planned_at).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
    
    return (
      <View style={styles.mealItem}>
        <Text style={styles.mealText}>Repas à: {mealTime}</Text>
        <Text style={styles.mealText}>Invités: {item.guests}</Text>
        
        {item.recipes && item.recipes.length > 0 ? (
          item.recipes.map((recipe) => (
            <TouchableOpacity
            onPress={() => {
              navigation.navigate("SingleRecipe", { id: recipe.id });
              onClose();
            }}
            >
              <Text style={styles.recipeText}>- {recipe.name}</Text>
            </TouchableOpacity>
          ))
        ) : (
          <Text style={styles.recipeText}>Pas de recette spécifiée</Text>
        )}
      </View>
    );
  };

  return (
    <Modal visible={visible} transparent={true} animationType="slide">
      <View style={styles.modalContainer}>
        <View style={styles.modalContent}>
          <Text style={styles.modalTitle}>Repas prévus le {selectedDate}</Text>
          {meals.length > 0 ? (
            <FlatList
              data={meals}
              keyExtractor={(item) => item.id.toString()}
              renderItem={renderMealItem}
            />
          ) : (
            <Text style={styles.noMealsText}>Aucun repas prévu ce jour-ci</Text>
          )}
          <TouchableOpacity onPress={onClose} style={styles.closeButton}>
            <Text style={styles.closeButtonText}>Fermer</Text>
          </TouchableOpacity>
          <TouchableOpacity onPress={() => setAddMealVisible(true)} style={styles.addButton}>
            <Text style={styles.addButtonText}>Planifier un repas</Text>
          </TouchableOpacity>
          {addMealVisible && (
            <AddMealForm
              visible={addMealVisible}
              onClose={() => setAddMealVisible(false)}
              selectedDate={selectedDate}
              onMealAdded={onClose} // This will close both modals after adding a new meal
            />
          )}
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
  mealItem: {
    padding: 10,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  mealText: {
    fontSize: 16,
  },
  noMealsText: {
    fontSize: 16,
    color: 'gray',
    textAlign: 'center',
    marginVertical: 20,
  },
  closeButton: {
    marginTop: 20,
    backgroundColor: 'blue',
    paddingVertical: 10,
    borderRadius: 5,
    alignItems: 'center',
  },
  closeButtonText: {
    color: 'white',
    fontSize: 16,
  },
  recipeText: {
    fontSize: 20,
    color: 'blue'
  },
  addButton: {
    marginTop: 10,
    padding: 10,
    backgroundColor: '#4CAF50',
    borderRadius: 5,
    alignItems: 'center',
  },
  addButtonText: {
    color: '#fff',
    fontWeight: 'bold',
  }
});


export default DateInfos;
