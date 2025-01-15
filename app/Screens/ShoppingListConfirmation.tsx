import React from 'react';
import { View, Text, StyleSheet, TouchableOpacity, TextInput } from 'react-native';
import { ComputeShoppingListInput, client } from '../Service/Api';

interface ShoppingListConfirmationProps {
  from: string;
  to: string;
  onClose: () => void;
  listName: string;
  onListNameChange: (name: string) => void;
}

const ShoppingListConfirmation: React.FC<ShoppingListConfirmationProps> = ({ from, to, onClose, listName, onListNameChange }) => {

    const handleGeneratePress = async () => {
        const user_id = await client.getUserIdHelper();
        const shoppingListInput: ComputeShoppingListInput = {
          user_id: user_id,
          from,
          to,
          name: listName || "Entrez un nom de liste",
        };
        console.log(shoppingListInput);

        try {
          await client.computeShoppingList(shoppingListInput);
          console.log("Shopping list computed successfully!");
        } catch (error) {
          console.error("Error computing shopping list:", error);
        } finally {
          onClose();
        }
    };

  return (
    <View style={styles.modalContainer}>
      <Text style={styles.title}>Confirm Date Range</Text>
      <Text style={styles.dateText}>Start Date: {from}</Text>
      <Text style={styles.dateText}>End Date: {to}</Text>


      <TextInput
        style={styles.input}
        placeholder="Enter list name"
        value={listName}
        onChangeText={onListNameChange}
      />

      <View style={styles.buttonContainer}>
        <TouchableOpacity onPress={onClose} style={styles.cancelButton}>
          <Text style={styles.buttonText}>Cancel</Text>
        </TouchableOpacity>


        <TouchableOpacity
          onPress={handleGeneratePress}
          style={[styles.generateButton, { opacity: listName ? 1 : 0.5 }]}
          disabled={!listName}
        >
          <Text style={styles.buttonText}>Generate</Text>
        </TouchableOpacity>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  modalContainer: {
    padding: 20,
    backgroundColor: 'white',
    borderRadius: 10,
    alignItems: 'center',
  },
  title: {
    fontSize: 20,
    fontWeight: 'bold',
    marginBottom: 15,
  },
  dateText: {
    fontSize: 16,
    marginVertical: 5,
  },
  buttonContainer: {
    flexDirection: 'row',
    marginTop: 20,
  },
  cancelButton: {
    backgroundColor: 'red',
    paddingVertical: 10,
    paddingHorizontal: 20,
    borderRadius: 5,
    marginRight: 10,
  },
  generateButton: {
    backgroundColor: 'green',
    paddingVertical: 10,
    paddingHorizontal: 20,
    borderRadius: 5,
  },
  buttonText: {
    color: 'white',
    fontWeight: 'bold',
  },
  input: {
    height: 40,
    borderColor: 'gray',
    borderWidth: 1,
    borderRadius: 5,
    paddingHorizontal: 10,
    marginBottom: 20,
    width: '100%',
  },
});

export default ShoppingListConfirmation;
