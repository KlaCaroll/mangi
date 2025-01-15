import AsyncStorage from "@react-native-async-storage/async-storage";
import { jwtDecode } from "jwt-decode";

interface Ingredient {
  id: number;
  name: string;
  quantity: number;
  unit: string;
}

export interface ErrorFormat {
  code: string;
  err: string;
}

export interface Recipe {
  id: number;
  name: string;
  owner_id: number;
  preparation_time?: number;
  total_time?: number;
  description: string;
  is_public: boolean;
  filter?: string[];
  ustensils?: string[];
  ingredients: Ingredient[];
}

export interface Meal {
  id: number;
  planned_at: string;
  guests: number;
  recipes: Recipe[] | null;
}

export interface MealInput {
  planned_at: string; // ISO date string format
  guests: number;
  recipe_id: number;
}

export interface MealOutput {
  id: number;
  err?: string;
}

export interface RecipeOutput {
  recipe: Recipe;
}

export interface FetchRecipesInput {
  page: number;
  name: string;
  preference: boolean;
}

export interface RecipesOutput {
  recipes: Recipe[];
  err?: string;
  code?: string;
}

export interface RecipesPageOutput {
  page: number;
  total: number;
  recipes: Recipe[];
}

export interface RegisterInput {
  email: string;
  password: string;
  name: string;
}

export interface RegisterOutput {
  id: number;
  err?: string;
}

export interface LoginInput {
  email: string;
  password: string;
}

export interface LoginOutput {
  id: number;
  token: string;
  err?: string;
}
interface CreateRecipeInput {
  name: string;
  preparation_time?: number;
  total_time?: number;
  description: string;
  is_public: boolean;
  categories?: { category_id: number }[];
  ustensils?: { ustensil_id: number }[];
  ingredients: {
    id: number;
    quantity: number;
    unit: string;
  }[];
}

interface CreateRecipeOutput {
  id: number;
}
export interface ShowUserOutput {
  id: number;
  name: string;
  email: string;
  ustensils: Ustensil[];
  preferences: Preference[];
  err?: string;
}

export interface Ustensil {
  ustensil_id: number;
  ustensil_name: string;
}

export interface Preference {
  preference_id: number;
  preference_name: string;
}

export interface ComputeShoppingListInput {
  user_id: number;
  from: string;
  to: string;
  name: string;
}

export interface FoodItem {
  name: string;
  unit: string;
  quantity: number;
}
export interface ShoppingList {
  name: string;
  user_id: number;
  from: string;
  to: string;
  home_id: number;
  items: FoodItem[];
}

export interface FetchShoppingListInput {
  from: string;
  to: string;
}

export interface Home {
  id: number;
  owner_id: number;
  name: string;
  owner_name: string;
  members: { id: number; name: string }[];
}
export interface SingleHome {
  id: number;
  name: string;
  owner_id: number;
  owner_name: string;
  members: { id: number; name: string }[];
  err?: string;
}

export interface DeleteShoppingListInput {
  user_id: number;
  from: string;
  to: string;
  name: string | undefined;
}

export interface DeleteItemInput {
  from: string;
  to: string;
  name: string | undefined;
  items: { name: string }[];
}

export interface HomeInvitationInput {
  home_id: number;
  invitation_to: string;
}

export interface HomeInvitationOutput {
  id: number;
  name: string;
  owner_id: number;
  owner_name: string;
  members: { id: number; name: string }[];
  err?: string;
}

class Client {
  endpoint: string;
  isShoppingListNameTaken: any;

  constructor(endpoint: string) {
    this.endpoint = endpoint;
  }

  async getHeaders(needToken: Boolean): Promise<HeadersInit | undefined> {
    let headers: HeadersInit | undefined = {
      "Content-Type": "application/json",
      "ngrok-skip-browser-warning": "true",
    };
    if (needToken) {
      const token = await AsyncStorage.getItem("userToken");
      if (!token) {
        throw new Error("No token found");
      }
      headers["Authorization"] = token;
    }
    return headers;
  }

  async getUserIdHelper(): Promise<number> {
    const token = await AsyncStorage.getItem("userToken");
    if (!token) {
      throw new Error("No token found");
    }
    const decodedToken = jwtDecode<{ sub: number }>(token);
    const userId = decodedToken.sub;
    return userId;
  }

  async validateToken(): Promise<boolean> {
    const token = await AsyncStorage.getItem("userToken");
    if (!token) return false;

    try {
      const decoded = jwtDecode<{ expired_at: string }>(token);
      const currentTime = new Date().getTime();
      const expiredAt = new Date(decoded.expired_at).getTime();

      return expiredAt > currentTime;
    } catch (error) {
      console.error("Failed to decode token", error);
      return false;
    }
  }

  // Login
  async login(input: LoginInput): Promise<LoginOutput> {
    const response = await fetch(`${this.endpoint}/login`, {
      method: "POST",
      headers: await this.getHeaders(false),
      body: JSON.stringify(input),
    });

    if (!response.ok) {
      throw new Error("Failed to login");
    }

    const data: LoginOutput = await response.json();
    await AsyncStorage.setItem("userToken", data.token);
    return data;
  }

  // Fetch all recipes
  //name hasn't been handled for now, having issue with the promise and the empty field
  async fetchRecipesPage(input: FetchRecipesInput): Promise<RecipesPageOutput> {
    const params = new URLSearchParams({
      page: input.page.toString(),
      name: input.name,
      preference: input.preference ? "true" : "false",
    }).toString();

    const response = await fetch(`${this.endpoint}/recipes?${params}`, {
      method: "GET",
      headers: await this.getHeaders(true),
    });

    if (!response.ok) {
      throw new Error("Failed to fetch recipes");
    }

    return await response.json();
  }

  async fetchRecipes(
    input: Omit<FetchRecipesInput, "page">
  ): Promise<RecipesOutput> {
    const perPage = 10;
    let recipes: Recipe[] = [];
    let page = 1;
    while (true) {
      const res = await this.fetchRecipesPage({ ...input, page });
      recipes.push(...res.recipes);
      if (res.recipes.length < perPage) {
        break;
      }
      page += 1;
    }
    return {
      recipes,
    };
  }

  async fetchRecipeById(id: number, guests: number): Promise<Recipe> {
    const response = await fetch(
      `${this.endpoint}/recipe?id=${id}&guests=${guests}`,
      {
        method: "GET",
        headers: await this.getHeaders(true),
      }
    );

    if (!response.ok) {
      throw new Error("Failed to fetch recipe");
    }
    return response.json();
  }

  async register(input: RegisterInput): Promise<RegisterOutput> {
    const response = await fetch(`${this.endpoint}/register`, {
      method: "POST",
      headers: await this.getHeaders(false),
      body: JSON.stringify(input),
    });

    if (!response.ok) {
      throw new Error("Failed to register");
    }
    return response.json();
  }

  async showProfile(id: number): Promise<ShowUserOutput> {
    const queryString = new URLSearchParams({
      id: id.toString(),
    }).toString();
    const response = await fetch(`${this.endpoint}/user?${queryString}`, {
      method: "GET",
      headers: await this.getHeaders(true),
    });

    if (!response.ok) {
      throw new Error("Failed to fetch user profile");
    }
    return response.json();
  }

  async fetchAllMeals() {
    const response = await fetch(`${this.endpoint}/meals`, {
      method: "GET",
      headers: await this.getHeaders(true),
    });

    if (!response.ok) {
      throw new Error("Failed to fetch Meals");
    }
    const data = await response.json();
    if (!data) {
      return [];
    }
    return data.meals;
  }

  async fetchMeals(from: string, to: string) {
    const response = await fetch(
      `${this.endpoint}/meals?from=${from}&to=${to}`,
      {
        method: "GET",
        headers: await this.getHeaders(true),
      }
    );

    if (!response.ok) {
      throw new Error("Failed to fetch Meals");
    }
    const data = await response.json();
    if (!data) {
      return [];
    }
    return data.meals;
  }

  createMeal = async (input: MealInput): Promise<MealOutput> => {
    const token = await AsyncStorage.getItem("userToken");
    if (!token) {
      throw new Error("No token found");
    }

    return fetch(`${this.endpoint}/meal`, {
      method: "POST",
      headers: {
        Authorization: `${token}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(input),
    })
      .then((res) => {
        if (!res.ok) {
          throw new Error("Failed to create meal");
        }
        return res.json();
      })
      .then((data) => {
        return data as MealOutput;
      });
  };

  async computeShoppingList(input: ComputeShoppingListInput) {
    input.from = `${input.from}T00:00:00.000Z`;
    input.to = `${input.to}T00:00:00.000Z`;
    input.name = `${input.name}`;
    const response = await fetch(`${this.endpoint}/compute-shopping-list`, {
      method: "POST",
      headers: await this.getHeaders(true),
      body: JSON.stringify(input),
    });
    const data = await response.json();
    if (data.err == "shopping list name empty") {
      throw new Error(" shopping list name empty");
    }

    if (!response.ok) {
      throw new Error("Failed to generate shopping list");
    }
    return response.json();
  }

  async fetchShoppingLists(): Promise<ShoppingList[]> {
    const response = await fetch(`${this.endpoint}/shopping-lists`, {
      method: "GET",
      headers: await this.getHeaders(true),
    });
    if (!response.ok) {
      throw new Error("Failed to generate shopping lists");
    }
    const data = await response.json();
    return data;
  }

  async fetchShoppingListByName(
    input: FetchShoppingListInput
  ): Promise<ShoppingList> {
    const response = await fetch(`${this.endpoint}/shopping-list`, {
      method: "POST",
      headers: await this.getHeaders(true),
      body: JSON.stringify(input),
    });

    if (!response.ok) {
      throw new Error("Failed to find the corresponding shopping list");
    }
    const data = await response.json();
    return data;
  }

  async deleteShoppingList(input: DeleteShoppingListInput) {
    const response = await fetch(`${this.endpoint}/shopping-list/delete`, {
      method: "PUT",
      headers: await this.getHeaders(true),
      body: JSON.stringify(input),
    });
    if (!response.ok) {
      throw new Error("Failed to find the corresponding shopping list");
    }
    const data = await response.json();
    return data;
  }

  async updateShoppingList(input: ShoppingList): Promise<ShoppingList> {
    const response = await fetch(`${this.endpoint}/shopping-list`, {
      method: "PUT",
      headers: await this.getHeaders(true),
      body: JSON.stringify(input),
    });
    if (!response.ok) {
      throw new Error("Failed to find the corresponding shopping list");
    }
    const data = await response.json();
    return data;
  }

  async deleteItems(input: DeleteItemInput): Promise<ShoppingList> {
    const response = await fetch(
      `${this.endpoint}/shopping-list/delete-items`,
      {
        method: "PUT",
        headers: await this.getHeaders(true),
        body: JSON.stringify(input),
      }
    );
    if (!response.ok) {
      throw new Error("Failed to find the corresponding shopping list");
    }
    const data = await response.json();
    return data;
  }

  async fetchHomes(): Promise<Home[]> {
    const response = await fetch(`${this.endpoint}/homes`, {
      method: "GET",
      headers: await this.getHeaders(true),
    });

    if (!response.ok) {
      throw new Error("Failed to fetch homes");
    }
    return response.json();
  }

  async createHome(home_name: string): Promise<Home> {
    const response = await fetch(`${this.endpoint}/user/home/create`, {
      method: "POST",
      headers: await this.getHeaders(true),
      body: JSON.stringify({ home_name }),
    });

    if (!response.ok) {
      throw new Error("Failed to create home");
    }
    return response.json();
  }

  async fetchSingleHome(home_name: string): Promise<SingleHome> {
    const response = await fetch(`${this.endpoint}/home`, {
      method: "POST",
      headers: await this.getHeaders(true),
      body: JSON.stringify({ home_name }),
    });
    if (!response.ok) {
      throw new Error("Failed to fetch home");
    }
    return response.json();
  }

  async deleteHome(home_name: string): Promise<void> {
    const response = await fetch(`${this.endpoint}/user/home/delete`, {
      method: "PUT",
      headers: await this.getHeaders(true),
      body: JSON.stringify({ home_name }),
    });
    if (!response.ok) {
      throw new Error("Failed to delete home");
    }
    return response.json();
  }

  async inviteUserHome(
    input: HomeInvitationInput
  ): Promise<HomeInvitationOutput> {
    const response = await fetch(`${this.endpoint}/home/invitation`, {
      method: "POST",
      headers: await this.getHeaders(true),
      body: JSON.stringify(input),
    });
    if (!response.ok) {
      throw new Error("Failed to invite user to home");
    }
    return response.json();
  }

  //  /!\ not useful unless we handle smtp host
  // async acceptHomeInvitation(home_id: number): Promise<void> {
  //   const response = await fetch(`${this.endpoint}/home/invitation`, {
  //     method: "PUT",
  //     headers: await this.getHeaders(true),
  //     body: JSON.stringify({ home_id }),
  //   });
  //   if (!response.ok) {
  //     throw new Error("Failed to accept home invitation");
  //   }
  //   return response.json();
  // }
}

export const client = new Client("https://c236-45-149-155-124.ngrok-free.app"); // localhost
//export const client = new Client("http://staging.mangi.local"); //staging
