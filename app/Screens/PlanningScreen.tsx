import React, { useState } from 'react';
import { SafeAreaView, ActivityIndicator, Alert, StyleSheet } from 'react-native';
import CalendarPicker from '../Components/Specific/CalendarPicker';
import DateInfos from '../Components/Specific/DateInfos';
import { client, Meal } from '../Service/Api';
import { useFocusEffect } from '@react-navigation/native';

export const PlanningScreen: React.FC<{ navigation: any }> = ({ navigation }) => {
  const [meals, setMeals] = useState<Meal[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [isDateInfosVisible, setDateInfosVisible] = useState(false);
  const [selectedDate, setSelectedDate] = useState<string>('');

  const fetchMeals = async () => {
    try {
    setLoading(true);

      const tokenValid = await client.validateToken();
      if (!tokenValid) {
        Alert.alert("Your session has expired. Please log in again.");
        navigation.navigate("Auth");
        setLoading(false);
        return;
      }

      const mealsData = await client.fetchMeals(
        "2021-07-01T19:00:00.000Z",
        "2045-11-02T09:00:00.000Z"
      );
      if (mealsData.err) {
        throw new Error(`An error occurred while fetching meals: ${mealsData.err}`);
      }
      if (mealsData) {
        setMeals(mealsData);
      }
    }
    catch (error: unknown) {
      console.error(error);

      if (error instanceof Error) {
        Alert.alert(error.message);
      } else {
        Alert.alert("An unexpected error occurred.");
      }
    }
      finally {
      setLoading(false);
    }
  };

  const onMealAdded = async () => {
    await fetchMeals();
  };

  useFocusEffect(
    React.useCallback(() => {
      fetchMeals();
    }, [navigation])
  );

  const handleOpenDateInfos = (date: string) => {
    setSelectedDate(date);
    setDateInfosVisible(true);
  };

  if (loading) {
    return (
      <SafeAreaView style={styles.loadingContainer}>
        <ActivityIndicator size="large" color="blue" />
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView style={styles.container}>
      <CalendarPicker meals={meals} onDateSelect={handleOpenDateInfos} onMealAdded={onMealAdded} />
      <DateInfos
        visible={isDateInfosVisible}
        selectedDate={selectedDate}
        onClose={() => setDateInfosVisible(false)}
        meals={meals.filter(meal => meal.planned_at.startsWith(selectedDate))}
      />
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    paddingTop: 10,
    backgroundColor: "white",
  },
  loadingContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
});

export default PlanningScreen;
