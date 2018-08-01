using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using YamlDotNet.Serialization;

namespace VulnerableDotNet
{
    class Program
    {
        static void Main(string[] args)
        {
            var r = new StringReader(@"
                scalar: a scalar
                sequence:
                - one
                - two
            ");

            var deserializer = new DeserializerBuilder().Build();
            var yamlObject = deserializer.Deserialize(r);
        }
    }
}
