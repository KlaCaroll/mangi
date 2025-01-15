import { StackNavigationProp } from "@react-navigation/stack";
import { RouteProp } from "@react-navigation/native";
import { FetchShoppingListInput } from "../Service/Api";

export type RecipesStackParamList = {
  Recipes: undefined;
  SingleRecipe: { id: number };
};

export type ProfileToHomesStackParamList = {
  Profile: { id: number };
  Homes: { userName: string | undefined };
  SingleHome: { home_id: number; home_name: string; refresh?: boolean };
  CreateHome: { setRefresh: React.Dispatch<React.SetStateAction<boolean>> };
  Invite: { home_id: number | undefined };
};

export type ShoppingListStackParamList = {
  ShoppingListsScreen: undefined;
  SingleShoppingListScreen: { input: FetchShoppingListInput };
};

export type ShoppingListStackNavigationProp =
  StackNavigationProp<ShoppingListStackParamList>;
  
export type SingleShoppingListRouteProp = RouteProp<
  ShoppingListStackParamList,
  "SingleShoppingListScreen"
>;

export type RecipesStackNavigationProp =
  StackNavigationProp<RecipesStackParamList>;

export type SingleRecipeRouteProp = RouteProp<
  RecipesStackParamList,
  "SingleRecipe"
>;

export type ProfileToHomesStackNavigationProp =
  StackNavigationProp<ProfileToHomesStackParamList>;
export type ProfileRouteProp = RouteProp<
  ProfileToHomesStackParamList,
  "Profile"
>;
export type HomesRouteProp = RouteProp<ProfileToHomesStackParamList, "Homes">;
export type SingleHomeRouteProp = RouteProp<
  ProfileToHomesStackParamList,
  "SingleHome"
>;
