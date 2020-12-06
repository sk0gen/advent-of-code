using System;
using System.Linq;
using System.Reflection;

namespace AOC_2020.Days
{

    public static class ReflectionHelper
    {
        //through reflection

//as a reusable method/function
        public static Type[] GetInheritedClasses(Type MyType)
        {
            //if you want the abstract classes drop the !TheType.IsAbstract but it is probably to instance so its a good idea to keep it.
            return Assembly.GetAssembly(MyType).GetTypes().Where(TheType =>
                TheType.IsClass && !TheType.IsAbstract && TheType.IsSubclassOf(MyType)).ToArray();
        }
    }

}