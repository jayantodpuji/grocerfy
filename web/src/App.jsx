import './App.css'

function App() {
  return (
    <div class="h-screen flex">
      <div class="w-1/2 bg-gray-200 p-8 flex items-center justify-center">
        <div>
          <h1 class="text-4xl font-bold">Welcome to Grocerfy</h1>
          <p class="mt-4 text-lg">A simple Grocery Todo List Application</p>
        </div>
      </div>
      <div class="w-1/2 p-8 flex items-center justify-center">
        <div>
          <h2 class="text-4xl font-bold mb-4">Login</h2>
          <form>
            <div class="mb-4">
              <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
              <input type="email" id="email" class="mt-1 block w-full p-3 border border-gray-300 rounded-md" placeholder="you@example.com"/>
            </div>
            <div class="mb-4">
              <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
              <input type="password" id="password" class="mt-1 block w-full p-3 border border-gray-300 rounded-md" placeholder="Enter your password"/>
            </div>
            <button type="submit" class="w-full py-3 px-4 bg-blue-600 text-white font-semibold rounded-md hover:bg-blue-700">Sign Up</button>
          </form>
        </div>
      </div>
    </div>
  )
}

export default App
