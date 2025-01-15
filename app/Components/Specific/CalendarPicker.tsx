import React, { useState } from 'react';
import { View, Text, StyleSheet, Button, Modal, TextInput } from 'react-native';
import { Calendar } from 'react-native-calendars';
import DateInfos from './DateInfos';
import { Meal } from '../../Service/Api';
import ShoppingListConfirmation from '../../Screens/ShoppingListConfirmation';

interface CalendarPickerProps {
  meals: Meal[];
  onDateSelect: (date: string) => void;
  onMealAdded: () => void;
}

const CalendarPicker: React.FC<CalendarPickerProps> = ({ meals, onMealAdded }) => {
  const [selectedDate, setSelectedDate] = useState<string>('');
  const [startDate, setStartDate] = useState<string>('');
  const [endDate, setEndDate] = useState<string>('');
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [mealsForSelectedDate, setMealsForSelectedDate] = useState<Meal[]>([]);
  const [isConfirmationVisible, setIsConfirmationVisible] = useState(false);
  const [listName, setListName] = useState<string>('');  // Ajout de l'état listName

  const onDayPress = (day: any) => {
    setSelectedDate(day.dateString);
    const selectedMeals = meals.filter((meal) => {
      const mealDate = meal.planned_at.split('T')[0];
      return mealDate === day.dateString;
    });
    setMealsForSelectedDate(selectedMeals);
    setIsModalVisible(true);
  };

  const onDayLongPress = (day: any) => {
    if (!startDate || (startDate && endDate)) {
      setStartDate(day.dateString);
      setEndDate("");
    } else if (startDate && !endDate) {
      setEndDate(day.dateString);
      setIsConfirmationVisible(true);
    }
  };

  const handleCloseModal = () => {
    setIsModalVisible(false);
    onMealAdded();
  };

  const handleCloseConfirmation = () => {
    setIsConfirmationVisible(false);
    setStartDate('');
    setEndDate('');
  };

  return (
    <View style={styles.container}>
      <Calendar
        onDayPress={onDayPress}
        onDayLongPress={onDayLongPress}
        markedDates={{
          [startDate || '']: { selected: true, selectedColor: 'blue' },
          [endDate || '']: { selected: true, selectedColor: 'blue' },
        }}
        style={styles.calendar}
      />

      <DateInfos
        visible={isModalVisible}
        onClose={handleCloseModal}
        selectedDate={selectedDate}
        meals={mealsForSelectedDate}
      />

      <Modal visible={isConfirmationVisible} transparent={true} animationType="slide">
        <View style={styles.modalOverlay}>
          <View style={styles.modalContent}>
            {/* Pass listName to ShoppingListConfirmation */}
            <ShoppingListConfirmation
              from={startDate}
              to={endDate}
              onClose={handleCloseConfirmation}
              listName={listName}  // Pass the listName prop here
              onListNameChange={setListName}  // Handler to update listName
            />
          </View>
        </View>
      </Modal>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
  },
  calendar: {
    borderWidth: 1,
    borderColor: 'gray',
    borderRadius: 10,
  },

  modalOverlay: {
    flex: 1,
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
    justifyContent: 'center',
    alignItems: 'center',
  },
  modalContent: {
    width: '80%',
    backgroundColor: 'white',
    padding: 20,
    borderRadius: 10,
  },
});

export default CalendarPicker;
