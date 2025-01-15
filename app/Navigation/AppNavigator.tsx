import React from "react";
import { StyleSheet } from "react-native";
import { NavigationContainer } from "@react-navigation/native";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { RecipesScreen } from "../Screens/RecipesScreen";
import { LoginScreen } from "../Screens/LoginScreen";
import { RegisterScreen } from "../Screens/RegisterScreen";
import { HomeScreen } from "../Screens/HomeScreen";
import { createStackNavigator } from "@react-navigation/stack";
import { createDrawerNavigator } from "@react-navigation/drawer";
import { SingleRecipeScreen } from "../Screens/SingleRecipeScreen";
import SingleShoppingListScreen from "../Screens/SingleShoppingListScreen";
import { ProfileScreen } from "../Screens/ProfileScreen";
import MaterialIcons from 'react-native-vector-icons/MaterialIcons';
import { AuthScreen } from "../Screens/AuthScreen";
import "react-native-gesture-handler";
import { PlanningScreen } from "../Screens/PlanningScreen";
import ShoppingListsScreen from "../Screens/ShoppingListsScreen";
import { FetchHomesScreen } from "../Screens/FetchHomesScreen";
import { SingleHomeScreen } from "../Screens/SingleHomeScreen";
import { CreateHomeScreen } from "../Screens/CreateHomeScreen";
import { InviteUserScreen } from "../Screens/InviteUserScreen";


const Tab = createBottomTabNavigator();
const Drawer = createDrawerNavigator();
const Stack = createStackNavigator();


function RecipesStackNavigator() {
  return (
    <Stack.Navigator initialRouteName="Recipes">
      <Stack.Screen
        name="Recipes"
        component={RecipesScreen}
        options={{ headerShown: false }}
      />
      <Stack.Screen
        name="SingleRecipe"
        component={SingleRecipeScreen}
        options={{title: "Ma recette" }}
      />
      <Stack.Screen name="ShoppingLists" component={ShoppingListsScreen} />
      <Stack.Screen
        name="SingleShoppingListScreen"
        component={SingleShoppingListScreen}
        options={{ title: "Shopping List Details" }}
      />
    </Stack.Navigator>
  );
}

function ProfileToHomesStackNavigator() {
  return (
    <Stack.Navigator initialRouteName="Profile">
      <Stack.Screen
        name="Profile"
        component={ProfileScreen}
        options={{ headerShown: false }}
      />
      <Stack.Screen
        name="Homes"
        component={FetchHomesScreen}
        options={{ headerShown: false }}
      />
      <Stack.Screen name="SingleHome" component={SingleHomeScreen} />
      <Stack.Screen name="CreateHome" component={CreateHomeScreen} />
      <Stack.Screen name="Invite" component={InviteUserScreen} />
    </Stack.Navigator>
  );
}


function ShoppingListStackNavigator() {
  return (
    <Stack.Navigator initialRouteName="ShoppingListsScreen">
      <Stack.Screen
        name="ShoppingListsScreen"
        component={ShoppingListsScreen}
        options={{ title: "Shopping Lists",
          headerShown : false,

         }}


      />
      <Stack.Screen
        name="SingleShoppingListScreen"
        component={SingleShoppingListScreen}
        options={{ title: "Shopping List Details",


         }}
      />
    </Stack.Navigator>
  );
}

export function TabNavigator() {
  return (
    <Tab.Navigator
      screenOptions={{
        tabBarActiveTintColor: "#FFCC00",
        tabBarInactiveTintColor: "black",
        tabBarStyle: {
          backgroundColor: "#FAF4EA",
          borderTopLeftRadius: 32,
          borderTopRightRadius: 32,
          height: 80,
          position: "absolute",
          overflow: "hidden",
          borderTopWidth: 0,
        },
      }}
    >
      <Tab.Screen
        name="Home Page"
        component={HomeScreen}
        options={{
          tabBarIcon: ({ color, size }) => (
            <MaterialIcons name="home" color={color} size={size} />
          ),
          headerShown: false,
        }}
      />
      <Tab.Screen
        name="Planning"
        component={PlanningScreen}
        options={{
          tabBarIcon: ({ color, size }) => (
            <MaterialIcons name="calendar-today" color={color} size={size} />
          ),
          headerShown: false,
        }}
      />
      <Tab.Screen
        name="Lists"
        component={ShoppingListStackNavigator}
        options={{
          tabBarIcon: ({ color, size }) => (
            <MaterialIcons name="shopping-cart" color={color} size={size} />
          ),
          headerShown: false,
        }}
      />
      <Tab.Screen
        name="Profile"
        component={ProfileToHomesStackNavigator}
        options={{
          tabBarIcon: ({ color, size }) => (
            <MaterialIcons name="person" color={color} size={size} />
          ),

          headerShown: false,
        }}
      />
      <Tab.Screen
        name="Recettes"
        component={RecipesStackNavigator}
        options={{
          tabBarIcon: ({ color, size }) => (
            <MaterialIcons name="restaurant" color={color} size={size} />
          ),
          headerShown: false,
        }}
      />
    </Tab.Navigator>
  );
}

function DrawerNavigator() {
  return (
    <Drawer.Navigator initialRouteName="Home">
      <Drawer.Screen name="Home" component={TabNavigator} />
      <Drawer.Screen name="Recettes" component={RecipesStackNavigator} />
      <Drawer.Screen name="Register" component={TabNavigator} />
      <Drawer.Screen name="Login" component={TabNavigator} />
    </Drawer.Navigator>
  );
}

function StackNavigator() {
  return (
    <Stack.Navigator initialRouteName="Login">
      <Stack.Screen name="Login" component={LoginScreen} />
      <Stack.Screen name="Planning" component={PlanningScreen} />
    </Stack.Navigator>
  );
}


export default function AppNavigator() {
  return (
    <NavigationContainer>
      <Stack.Navigator initialRouteName="Auth">
        <Stack.Screen
          name="Auth"
          component={AuthScreen}
          options={{ headerShown: false }}
        />
        <Stack.Screen
          name="Main"
          component={TabNavigator}
          options={{ headerShown: false }}
        />
        <Stack.Screen
          name="Register"
          component={RegisterScreen}
          options={{ headerShown: false }}
        />
        <Stack.Screen
          name="Login"
          component={LoginScreen}
          options={{ headerShown: false }}
        />
      </Stack.Navigator>

      {/* <TabNavigator /> */}
    </NavigationContainer>
  );
}

const styles = StyleSheet.create({
  homeContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
  title: {
    fontSize: 24,
    fontWeight: "bold",
  },
});
