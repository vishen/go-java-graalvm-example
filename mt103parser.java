import org.graalvm.nativeimage.IsolateThread;
import org.graalvm.nativeimage.c.function.CEntryPoint;
import org.graalvm.nativeimage.c.type.CCharPointer;
import org.graalvm.nativeimage.c.type.CCharPointerPointer;
import org.graalvm.nativeimage.c.type.CTypeConversion;

import com.prowidesoftware.swift.model.SwiftMessage;
import com.prowidesoftware.swift.model.mt.mt1xx.MT103;
import com.prowidesoftware.swift.io.parser.SwiftParser;

public class MT103Parser {
    @CEntryPoint(name = "mt103_prowide_example")
    private static CCharPointer mt103Parser(IsolateThread thread) {
		try {
			// https://github.com/prowide/prowide-core-examples/blob/master/src/main/java/com/prowidesoftware/swift/samples/core/MessageToJsonExample.java
			String fin = "{1:F01BICFOOYYAXXX8683497519}{2:O1031535051028ESPBESMMAXXX54237522470510281535N}{3:{113:ROMF}{108:0510280182794665}{119:STP}}{4:\n" +
					":20:0061350113089908\n" +
					":13C:/RNCTIME/1534+0000\n" +
					":23B:CRED\n" +
					":23E:SDVA\n" +
					":32A:061028EUR100000,\n" +
					":33B:EUR100000,\n" +
					":50K:/12345678\n" +
					"AGENTES DE BOLSA FOO AGENCIA\n" +
					"AV XXXXX 123 BIS 9 PL\n" +
					"12345 BARCELONA\n" +
					":52A:/2337\n" +
					"FOOAESMMXXX\n" +
					":53A:FOOAESMMXXX\n" +
					":57A:BICFOOYYXXX\n" +
					":59:/ES0123456789012345671234\n" +
					"FOO AGENTES DE BOLSA ASOC\n" +
					":71A:OUR\n" +
					":72:/BNF/TRANSF. BCO. FOO\n" +
					"-}{5:{MAC:88B4F929}{CHK:22EF370A4073}}";
			SwiftParser parser = new SwiftParser(fin);
			SwiftMessage mt = parser.message();
			System.out.println("Sender: "+mt.getSender());
			System.out.println("Receiver: "+mt.getReceiver());
			System.out.println(mt.toJson());
			return CTypeConversion.toCString(mt.toJson()).get();
		} catch (Exception e) {
			System.out.format("error: %s\n", e.getMessage());
			return CTypeConversion.toCString("error").get();
		}
    }

    @CEntryPoint(name = "mt103_parser")
    private static CCharPointer mt103Parser(IsolateThread thread, CCharPointer cIncoming) {
		try {
			String incoming = CTypeConversion.toJavaString(cIncoming);
			MT103 mt103 = MT103.parse(incoming);
			System.out.println("Sender: "+mt103.getSender());
			System.out.println("Receiver: "+mt103.getReceiver());
			System.out.println(mt103);
			System.out.println(mt103.getSwiftMessage().toJson());
			return CTypeConversion.toCString(mt103.toJson()).get();
		} catch (Exception e) {
			System.out.format("error: %s\n", e.getMessage());
			return CTypeConversion.toCString("error").get();
		}
    }
}
