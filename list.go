package mime2extension

type BidirectionalMap struct {
	mimeToExt map[string][]string
	extToMime map[string]string
}

var bdMap *BidirectionalMap

//init when load lib
func init() {
	mimeToExt := createMimeToExt()
	extToMime := createExtToMime(mimeToExt)
	bdMap = &BidirectionalMap {
		extToMime: extToMime,
		mimeToExt: mimeToExt,
	}
	
}

func createExtToMime(mimeToExt map[string][]string) map[string]string {
	extToMime := map[string]string {}
	for k, exts := range mimeToExt {
		for _, ext := range exts {
			extToMime[ext] = k
		}
	}
	return extToMime
}

func createMimeToExt() map[string][]string {
	mimeToExt := map[string][]string {
		 "application/vnd.lotus-1-2-3": []string{"123"},
		 "text/vnd.in3d.3dml": []string{"3dml"},
		 "image/x-3ds": []string{"3ds"},
		 "video/3gpp2": []string{"3g2"},
		 "video/3gpp": []string{"3gp"},
		 "application/x-7z-compressed": []string{"7z"},
		 "application/x-authorware-bin": []string{"aab", "u32", "vox", "x32"},
		 "audio/x-aac": []string{"aac"},
		 "application/x-authorware-map": []string{"aam"},
		 "application/x-authorware-seg": []string{"aas"},
		 "application/x-abiword": []string{"abw"},
		 "application/pkix-attr-cert": []string{"ac"},
		 "application/vnd.americandynamics.acc": []string{"acc"},
		 "application/x-ace-compressed": []string{"ace"},
		 "application/vnd.acucobol": []string{"acu"},
		 "application/vnd.acucorp": []string{"acutc", "atc"},
		 "audio/adpcm": []string{"adp"},
		 "application/vnd.audiograph": []string{"aep"},
		 "application/x-font-type1": []string{"afm", "pfa", "pfb", "pfm"},
		 "application/vnd.ibm.modcap": []string{"afp", "list3820", "listafp"},
		 "application/vnd.ahead.space": []string{"ahead"},
		 "application/postscript": []string{"ai", "eps", "ps"},
		 "audio/x-aiff": []string{"aif", "aifc", "aiff"},
		 "application/vnd.adobe.air-application-installer-package+zip": []string{"air"},
		 "application/vnd.dvb.ait": []string{"ait"},
		 "application/vnd.amiga.ami": []string{"ami"},
		 "application/vnd.android.package-archive": []string{"apk"},
		 "text/cache-manifest": []string{"appcache"},
		 "application/x-ms-application": []string{"application"},
		 "application/vnd.lotus-approach": []string{"apr"},
		 "application/x-freearc": []string{"arc"},
		 "application/pgp-signature": []string{"asc", "sig"},
		 "video/x-ms-asf": []string{"asf", "asx"},
		 "text/x-asm": []string{"asm", "s"},
		 "application/vnd.accpac.simply.aso": []string{"aso"},
		 "application/atom+xml": []string{"atom"},
		 "application/atomcat+xml": []string{"atomcat"},
		 "application/atomsvc+xml": []string{"atomsvc"},
		 "application/vnd.antix.game-component": []string{"atx"},
		 "audio/basic": []string{"au", "snd"},
		 "video/x-msvideo": []string{"avi"},
		 "application/applixware": []string{"aw"},
		 "application/vnd.airzip.filesecure.azf": []string{"azf"},
		 "application/vnd.airzip.filesecure.azs": []string{"azs"},
		 "application/vnd.amazon.ebook": []string{"azw"},
		 "application/x-msdownload": []string{"bat", "com", "dll", "exe", "msi"},
		 "application/x-bcpio": []string{"bcpio"},
		 "application/x-font-bdf": []string{"bdf"},
		 "application/vnd.syncml.dm+wbxml": []string{"bdm"},
		 "application/vnd.realvnc.bed": []string{"bed"},
		 "application/vnd.fujitsu.oasysprs": []string{"bh2"},
		 "application/x-blorb": []string{"blb", "blorb"},
		 "application/vnd.bmi": []string{"bmi"},
		 "image/bmp": []string{"bmp"},
		 "application/vnd.framemaker": []string{"book", "fm", "frame", "maker"},
		 "application/vnd.previewsystems.box": []string{"box"},
		 "application/x-bzip2": []string{"boz", "bz2"},
		 "image/prs.btif": []string{"btif"},
		 "application/x-bzip": []string{"bz"},
		 "application/vnd.cluetrust.cartomobile-config": []string{"c11amc"},
		 "application/vnd.cluetrust.cartomobile-config-pkg": []string{"c11amz"},
		 "application/vnd.clonk.c4group": []string{"c4d", "c4f", "c4g", "c4p", "c4u"},
		 "application/vnd.ms-cab-compressed": []string{"cab"},
		 "audio/x-caf": []string{"caf"},
		 "application/vnd.tcpdump.pcap": []string{"cap", "dmp", "pcap"},
		 "application/vnd.curl.car": []string{"car"},
		 "application/vnd.ms-pki.seccat": []string{"cat"},
		 "application/x-cbr": []string{"cb7", "cba", "cbr", "cbt", "cbz"},
		 "application/x-director": []string{"cct", "cst", "cxt", "dcr", "dir", "dxr", "fgd", "swa", "w3d"},
		 "text/x-c": []string{"cc", "cpp", "c", "cxx", "dic", "hh", "h"},
		 "application/ccxml+xml": []string{"ccxml"},
		 "application/vnd.contact.cmsg": []string{"cdbcmsg"},
		 "application/x-netcdf": []string{"cdf", "nc"},
		 "application/vnd.mediastation.cdkey": []string{"cdkey"},
		 "application/cdmi-capability": []string{"cdmia"},
		 "application/cdmi-container": []string{"cdmic"},
		 "application/cdmi-domain": []string{"cdmid"},
		 "application/cdmi-object": []string{"cdmio"},
		 "application/cdmi-queue": []string{"cdmiq"},
		 "chemical/x-cdx": []string{"cdx"},
		 "application/vnd.chemdraw+xml": []string{"cdxml"},
		 "application/vnd.cinderella": []string{"cdy"},
		 "application/pkix-cert": []string{"cer"},
		 "application/x-cfs-compressed": []string{"cfs"},
		 "image/cgm": []string{"cgm"},
		 "application/x-chat": []string{"chat"},
		 "application/vnd.ms-htmlhelp": []string{"chm"},
		 "application/vnd.kde.kchart": []string{"chrt"},
		 "chemical/x-cif": []string{"cif"},
		 "application/vnd.anser-web-certificate-issue-initiation": []string{"cii"},
		 "application/vnd.ms-artgalry": []string{"cil"},
		 "application/vnd.claymore": []string{"cla"},
		 "application/java-vm": []string{"class"},
		 "application/vnd.crick.clicker.keyboard": []string{"clkk"},
		 "application/vnd.crick.clicker.palette": []string{"clkp"},
		 "application/vnd.crick.clicker.template": []string{"clkt"},
		 "application/vnd.crick.clicker.wordbank": []string{"clkw"},
		 "application/vnd.crick.clicker": []string{"clkx"},
		 "application/x-msclip": []string{"clp"},
		 "application/vnd.cosmocaller": []string{"cmc"},
		 "chemical/x-cmdf": []string{"cmdf"},
		 "chemical/x-cml": []string{"cml"},
		 "application/vnd.yellowriver-custom-menu": []string{"cmp"},
		 "image/x-cmx": []string{"cmx"},
		 "application/vnd.rim.cod": []string{"cod"},
		 "text/plain": []string{"txt", "conf", "def", "in", "list", "log", "text"},
		 "application/x-cpio": []string{"cpio"},
		 "application/mac-compactpro": []string{"cpt"},
		 "application/x-mscardfile": []string{"crd"},
		 "application/pkix-crl": []string{"crl"},
		 "application/x-x509-ca-cert": []string{"crt", "der"},
		 "application/vnd.rig.cryptonote": []string{"cryptonote"},
		 "application/x-csh": []string{"csh"},
		 "chemical/x-csml": []string{"csml"},
		 "application/vnd.commonspace": []string{"csp"},
		 "text/css": []string{"css"},
		 "text/csv": []string{"csv"},
		 "application/cu-seeme": []string{"cu"},
		 "text/vnd.curl": []string{"curl"},
		 "application/prs.cww": []string{"cww"},
		 "model/vnd.collada+xml": []string{"dae"},
		 "application/vnd.mobius.daf": []string{"daf"},
		 "application/vnd.dart": []string{"dart"},
		 "application/vnd.fdsn.seed": []string{"dataless", "seed"},
		 "application/davmount+xml": []string{"davmount"},
		 "application/docbook+xml": []string{"dbk"},
		 "text/vnd.curl.dcurl": []string{"dcurl"},
		 "application/vnd.oma.dd2+xml": []string{"dd2"},
		 "application/vnd.fujixerox.ddd": []string{"ddd"},
		 "application/x-debian-package": []string{"deb", "udeb"},
		 "application/octet-stream": []string{"bin", "bpk", "deploy", "dist", "distz", "dms", "dump", "elc", "lrf", "mar", "pkg", "so"},
		 "application/vnd.dreamfactory": []string{"dfac"},
		 "application/x-dgc-compressed": []string{"dgc"},
		 "application/vnd.mobius.dis": []string{"dis"},
		 "image/vnd.djvu": []string{"djv", "djvu"},
		 "application/x-apple-diskimage": []string{"dmg"},
		 "application/vnd.dna": []string{"dna"},
		 "application/msword": []string{"doc", "dot"},
		 "application/vnd.ms-word.document.macroenabled.12": []string{"docm"},
		 "application/vnd.openxmlformats-officedocument.wordprocessingml.document": []string{"docx"},
		 "application/vnd.ms-word.template.macroenabled.12": []string{"dotm"},
		 "application/vnd.openxmlformats-officedocument.wordprocessingml.template": []string{"dotx"},
		 "application/vnd.osgi.dp": []string{"dp"},
		 "application/vnd.dpgraph": []string{"dpg"},
		 "audio/vnd.dra": []string{"dra"},
		 "text/prs.lines.tag": []string{"dsc"},
		 "application/dssc+der": []string{"dssc"},
		 "application/x-dtbook+xml": []string{"dtb"},
		 "application/xml-dtd": []string{"dtd"},
		 "audio/vnd.dts": []string{"dts"},
		 "audio/vnd.dts.hd": []string{"dtshd"},
		 "video/vnd.dvb.file": []string{"dvb"},
		 "application/x-dvi": []string{"dvi"},
		 "model/vnd.dwf": []string{"dwf"},
		 "image/vnd.dwg": []string{"dwg"},
		 "image/vnd.dxf": []string{"dxf"},
		 "application/vnd.spotfire.dxp": []string{"dxp"},
		 "audio/vnd.nuera.ecelp4800": []string{"ecelp4800"},
		 "audio/vnd.nuera.ecelp7470": []string{"ecelp7470"},
		 "audio/vnd.nuera.ecelp9600": []string{"ecelp9600"},
		 "application/ecmascript": []string{"ecma"},
		 "application/vnd.novadigm.edm": []string{"edm"},
		 "application/vnd.novadigm.edx": []string{"edx"},
		 "application/vnd.picsel": []string{"efif"},
		 "application/vnd.pg.osasli": []string{"ei6"},
		 "application/x-msmetafile": []string{"emf", "emz", "wmf", "wmz"},
		 "message/rfc822": []string{"eml", "mime"},
		 "application/emma+xml": []string{"emma"},
		 "audio/vnd.digital-winds": []string{"eol"},
		 "application/vnd.ms-fontobject": []string{"eot"},
		 "application/epub+zip": []string{"epub"},
		 "application/vnd.eszigno3+xml": []string{"es3", "et3"},
		 "application/vnd.osgi.subsystem": []string{"esa"},
		 "application/vnd.epson.esf": []string{"esf"},
		 "text/x-setext": []string{"etx"},
		 "application/x-eva": []string{"eva"},
		 "application/x-envoy": []string{"evy"},
		 "application/exi": []string{"exi"},
		 "application/vnd.novadigm.ext": []string{"ext"},
		 "application/vnd.ezpix-album": []string{"ez2"},
		 "application/vnd.ezpix-package": []string{"ez3"},
		 "application/andrew-inset": []string{"ez"},
		 "video/x-f4v": []string{"f4v"},
		 "text/x-fortran": []string{"f77", "f90", "for", "f"},
		 "image/vnd.fastbidsheet": []string{"fbs"},
		 "application/vnd.adobe.formscentral.fcdt": []string{"fcdt"},
		 "application/vnd.isac.fcs": []string{"fcs"},
		 "application/vnd.fdf": []string{"fdf"},
		 "application/vnd.denovo.fcselayout-link": []string{"fe_launch"},
		 "application/vnd.fujitsu.oasysgp": []string{"fg5"},
		 "image/x-freehand": []string{"fh4", "fh5", "fh7", "fhc", "fh"},
		 "application/x-xfig": []string{"fig"},
		 "audio/x-flac": []string{"flac"},
		 "video/x-fli": []string{"fli"},
		 "application/vnd.micrografx.flo": []string{"flo"},
		 "video/x-flv": []string{"flv"},
		 "application/vnd.kde.kivio": []string{"flw"},
		 "text/vnd.fmi.flexstor": []string{"flx"},
		 "text/vnd.fly": []string{"fly"},
		 "application/vnd.frogans.fnc": []string{"fnc"},
		 "image/vnd.fpx": []string{"fpx"},
		 "application/vnd.fsc.weblaunch": []string{"fsc"},
		 "image/vnd.fst": []string{"fst"},
		 "application/vnd.fluxtime.clip": []string{"ftc"},
		 "application/vnd.anser-web-funds-transfer-initiation": []string{"fti"},
		 "video/vnd.fvt": []string{"fvt"},
		 "application/vnd.adobe.fxp": []string{"fxp", "fxpl"},
		 "application/vnd.fuzzysheet": []string{"fzs"},
		 "application/vnd.geoplan": []string{"g2w"},
		 "image/g3fax": []string{"g3"},
		 "application/vnd.geospace": []string{"g3w"},
		 "application/vnd.groove-account": []string{"gac"},
		 "application/x-tads": []string{"gam"},
		 "application/rpki-ghostbusters": []string{"gbr"},
		 "application/x-gca-compressed": []string{"gca"},
		 "model/vnd.gdl": []string{"gdl"},
		 "application/vnd.dynageo": []string{"geo"},
		 "application/vnd.geometry-explorer": []string{"gex", "gre"},
		 "application/vnd.geogebra.file": []string{"ggb"},
		 "application/vnd.geogebra.tool": []string{"ggt"},
		 "application/vnd.groove-help": []string{"ghf"},
		 "image/gif": []string{"gif"},
		 "application/vnd.groove-identity-message": []string{"gim"},
		 "application/gml+xml": []string{"gml"},
		 "application/vnd.gmx": []string{"gmx"},
		 "application/x-gnumeric": []string{"gnumeric"},
		 "application/vnd.flographit": []string{"gph"},
		 "application/gpx+xml": []string{"gpx"},
		 "application/vnd.grafeq": []string{"gqf", "gqs"},
		 "application/srgs": []string{"gram"},
		 "application/x-gramps-xml": []string{"gramps"},
		 "application/vnd.groove-injector": []string{"grv"},
		 "application/srgs+xml": []string{"grxml"},
		 "application/x-font-ghostscript": []string{"gsf"},
		 "application/x-gtar": []string{"gtar"},
		 "application/vnd.groove-tool-message": []string{"gtm"},
		 "model/vnd.gtw": []string{"gtw"},
		 "text/vnd.graphviz": []string{"gv"},
		 "application/gxf": []string{"gxf"},
		 "application/vnd.geonext": []string{"gxt"},
		 "video/h261": []string{"h261"},
		 "video/h263": []string{"h263"},
		 "video/h264": []string{"h264"},
		 "application/vnd.hal+xml": []string{"hal"},
		 "application/vnd.hbci": []string{"hbci"},
		 "application/x-hdf": []string{"hdf"},
		 "application/winhlp": []string{"hlp"},
		 "application/vnd.hp-hpgl": []string{"hpgl"},
		 "application/vnd.hp-hpid": []string{"hpid"},
		 "application/vnd.hp-hps": []string{"hps"},
		 "application/mac-binhex40": []string{"hqx"},
		 "application/vnd.kenameaapp": []string{"htke"},
		 "text/html": []string{"html", "htm"},
		 "application/vnd.yamaha.hv-dic": []string{"hvd"},
		 "application/vnd.yamaha.hv-voice": []string{"hvp"},
		 "application/vnd.yamaha.hv-script": []string{"hvs"},
		 "application/vnd.intergeo": []string{"i2g"},
		 "application/vnd.iccprofile": []string{"icc", "icm"},
		 "x-conference/x-cooltalk": []string{"ice"},
		 "image/x-icon": []string{"ico"},
		 "text/calendar": []string{"ics", "ifb"},
		 "image/ief": []string{"ief"},
		 "application/vnd.shana.informed.formdata": []string{"ifm"},
		 "model/iges": []string{"iges", "igs"},
		 "application/vnd.igloader": []string{"igl"},
		 "application/vnd.insors.igm": []string{"igm"},
		 "application/vnd.micrografx.igx": []string{"igx"},
		 "application/vnd.shana.informed.interchange": []string{"iif"},
		 "application/vnd.accpac.simply.imp": []string{"imp"},
		 "application/vnd.ms-ims": []string{"ims"},
		 "application/inkml+xml": []string{"ink", "inkml"},
		 "application/x-install-instructions": []string{"install"},
		 "application/vnd.astraea-software.iota": []string{"iota"},
		 "application/ipfix": []string{"ipfix"},
		 "application/vnd.shana.informed.package": []string{"ipk"},
		 "application/vnd.ibm.rights-management": []string{"irm"},
		 "application/vnd.irepository.package+xml": []string{"irp"},
		 "application/x-iso9660-image": []string{"iso"},
		 "application/vnd.shana.informed.formtemplate": []string{"itp"},
		 "application/vnd.immervision-ivp": []string{"ivp"},
		 "application/vnd.immervision-ivu": []string{"ivu"},
		 "text/vnd.sun.j2me.app-descriptor": []string{"jad"},
		 "application/vnd.jam": []string{"jam"},
		 "application/java-archive": []string{"jar"},
		 "text/x-java-source": []string{"java"},
		 "application/vnd.jisp": []string{"jisp"},
		 "application/vnd.hp-jlyt": []string{"jlt"},
		 "application/x-java-jnlp-file": []string{"jnlp"},
		 "application/vnd.joost.joda-archive": []string{"joda"},
		 "image/jpeg": []string{"jpeg", "jpe", "jpg"},
		 "video/jpm": []string{"jpgm", "jpm"},
		 "video/jpeg": []string{"jpgv"},
		 "application/javascript": []string{"js"},
		 "application/json": []string{"json"},
		 "application/jsonml+json": []string{"jsonml"},
		 "audio/midi": []string{"kar", "mid", "midi", "rmi"},
		 "application/vnd.kde.karbon": []string{"karbon"},
		 "application/vnd.kde.kformula": []string{"kfo"},
		 "application/vnd.kidspiration": []string{"kia"},
		 "application/vnd.google-earth.kml+xml": []string{"kml"},
		 "application/vnd.google-earth.kmz": []string{"kmz"},
		 "application/vnd.kinar": []string{"kne", "knp"},
		 "application/vnd.kde.kontour": []string{"kon"},
		 "application/vnd.kde.kpresenter": []string{"kpr", "kpt"},
		 "application/vnd.ds-keypoint": []string{"kpxx"},
		 "application/vnd.kde.kspread": []string{"ksp"},
		 "application/vnd.kahootz": []string{"ktr", "ktz"},
		 "image/ktx": []string{"ktx"},
		 "application/vnd.kde.kword": []string{"kwd", "kwt"},
		 "application/vnd.las.las+xml": []string{"lasxml"},
		 "application/x-latex": []string{"latex"},
		 "application/vnd.llamagraphics.life-balance.desktop": []string{"lbd"},
		 "application/vnd.llamagraphics.life-balance.exchange+xml": []string{"lbe"},
		 "application/vnd.hhe.lesson-player": []string{"les"},
		 "application/x-lzh-compressed": []string{"lha"},
		 "application/vnd.route66.link66+xml": []string{"link66"},
		 "application/x-ms-shortcut": []string{"lnk"},
		 "application/lost+xml": []string{"lostxml"},
		 "application/vnd.ms-lrm": []string{"lrm"},
		 "application/vnd.frogans.ltf": []string{"ltf"},
		 "audio/vnd.lucent.voice": []string{"lvp"},
		 "application/vnd.lotus-wordpro": []string{"lwp"},
		 "application/x-msmediaview": []string{"m13", "m14", "mvb"},
		 "video/mpeg": []string{"mpeg", "m1v", "m2v", "mpe", "mpg"},
		 "application/mp21": []string{"m21", "mp21"},
		 "audio/mpeg": []string{"m2a", "m3a", "mp2a", "mp2", "mp3", "mpga"},
		 "application/vnd.apple.mpegurl": []string{"m3u8"},
		 "audio/x-mpegurl": []string{"m3u"},
		 "video/vnd.mpegurl": []string{"m4u", "mxu"},
		 "video/x-m4v": []string{"m4v"},
		 "application/mathematica": []string{"ma", "mb", "nb"},
		 "application/mads+xml": []string{"mads"},
		 "application/vnd.ecowin.chart": []string{"mag"},
		 "text/troff": []string{"man", "me", "ms", "roff", "tr", "t"},
		 "application/mathml+xml": []string{"mathml"},
		 "application/vnd.mobius.mbk": []string{"mbk"},
		 "application/mbox": []string{"mbox"},
		 "application/vnd.medcalcdata": []string{"mc1"},
		 "application/vnd.mcd": []string{"mcd"},
		 "text/vnd.curl.mcurl": []string{"mcurl"},
		 "application/x-msaccess": []string{"mdb"},
		 "image/vnd.ms-modi": []string{"mdi"},
		 "model/mesh": []string{"mesh", "msh", "silo"},
		 "application/metalink4+xml": []string{"meta4"},
		 "application/metalink+xml": []string{"metalink"},
		 "application/mets+xml": []string{"mets"},
		 "application/vnd.mfmp": []string{"mfm"},
		 "application/rpki-manifest": []string{"mft"},
		 "application/vnd.osgeo.mapguide.package": []string{"mgp"},
		 "application/vnd.proteus.magazine": []string{"mgz"},
		 "application/x-mie": []string{"mie"},
		 "application/vnd.mif": []string{"mif"},
		 "video/mj2": []string{"mj2", "mjp2"},
		 "video/x-matroska": []string{"mk3d", "mks", "mkv"},
		 "audio/x-matroska": []string{"mka"},
		 "application/vnd.dolby.mlp": []string{"mlp"},
		 "application/vnd.chipnuts.karaoke-mmd": []string{"mmd"},
		 "application/vnd.smaf": []string{"mmf"},
		 "image/vnd.fujixerox.edmics-mmr": []string{"mmr"},
		 "video/x-mng": []string{"mng"},
		 "application/x-msmoney": []string{"mny"},
		 "application/x-mobipocket-ebook": []string{"mobi", "prc"},
		 "application/mods+xml": []string{"mods"},
		 "video/x-sgi-movie": []string{"movie"},
		 "video/quicktime": []string{"mov", "qt"},
		 "audio/mp4": []string{"mp4a"},
		 "application/mp4": []string{"mp4s"},
		 "video/mp4": []string{"mp4", "mp4v", "mpg4"},
		 "application/vnd.mophun.certificate": []string{"mpc"},
		 "application/vnd.apple.installer+xml": []string{"mpkg"},
		 "application/vnd.blueice.multipass": []string{"mpm"},
		 "application/vnd.mophun.application": []string{"mpn"},
		 "application/vnd.ms-project": []string{"mpp", "mpt"},
		 "application/vnd.ibm.minipay": []string{"mpy"},
		 "application/vnd.mobius.mqy": []string{"mqy"},
		 "application/marc": []string{"mrc"},
		 "application/marcxml+xml": []string{"mrcx"},
		 "application/mediaservercontrol+xml": []string{"mscml"},
		 "application/vnd.fdsn.mseed": []string{"mseed"},
		 "application/vnd.mseq": []string{"mseq"},
		 "application/vnd.epson.msf": []string{"msf"},
		 "application/vnd.mobius.msl": []string{"msl"},
		 "application/vnd.muvee.style": []string{"msty"},
		 "model/vnd.mts": []string{"mts"},
		 "application/vnd.musician": []string{"mus"},
		 "application/vnd.recordare.musicxml+xml": []string{"musicxml"},
		 "application/vnd.mfer": []string{"mwf"},
		 "application/mxf": []string{"mxf"},
		 "application/vnd.recordare.musicxml": []string{"mxl"},
		 "application/xv+xml": []string{"xhvml", "xvm", "mxml"},
		 "application/vnd.triscape.mxs": []string{"mxs"},
		 "text/n3": []string{"n3"},
		 "application/vnd.wolfram.player": []string{"nbp"},
		 "application/x-dtbncx+xml": []string{"ncx"},
		 "text/x-nfo": []string{"nfo"},
		 "application/vnd.nokia.n-gage.symbian.install": []string{"n-gage"},
		 "application/vnd.nokia.n-gage.data": []string{"ngdat"},
		 "application/vnd.nitf": []string{"nitf", "ntf"},
		 "application/vnd.neurolanguage.nlu": []string{"nlu"},
		 "application/vnd.enliven": []string{"nml"},
		 "application/vnd.noblenet-directory": []string{"nnd"},
		 "application/vnd.noblenet-sealer": []string{"nns"},
		 "application/vnd.noblenet-web": []string{"nnw"},
		 "image/vnd.net-fpx": []string{"npx"},
		 "application/x-conference": []string{"nsc"},
		 "application/vnd.lotus-notes": []string{"nsf"},
		 "application/x-nzb": []string{"nzb"},
		 "application/vnd.fujitsu.oasys2": []string{"oa2"},
		 "application/vnd.fujitsu.oasys3": []string{"oa3"},
		 "application/vnd.fujitsu.oasys": []string{"oas"},
		 "application/x-msbinder": []string{"obd"},
		 "application/x-tgif": []string{"obj"},
		 "application/oda": []string{"oda"},
		 "application/vnd.oasis.opendocument.database": []string{"odb"},
		 "application/vnd.oasis.opendocument.chart": []string{"odc"},
		 "application/vnd.oasis.opendocument.formula": []string{"odf"},
		 "application/vnd.oasis.opendocument.formula-template": []string{"odft"},
		 "application/vnd.oasis.opendocument.graphics": []string{"odg"},
		 "application/vnd.oasis.opendocument.image": []string{"odi"},
		 "application/vnd.oasis.opendocument.text-master": []string{"odm"},
		 "application/vnd.oasis.opendocument.presentation": []string{"odp"},
		 "application/vnd.oasis.opendocument.spreadsheet": []string{"ods"},
		 "application/vnd.oasis.opendocument.text": []string{"odt"},
		 "audio/ogg": []string{"oga", "ogg", "spx"},
		 "video/ogg": []string{"ogv"},
		 "application/ogg": []string{"ogx"},
		 "application/omdoc+xml": []string{"omdoc"},
		 "application/onenote": []string{"onepkg", "onetmp", "onetoc2", "onetoc"},
		 "application/oebps-package+xml": []string{"opf"},
		 "text/x-opml": []string{"opml"},
		 "application/vnd.palm": []string{"oprc", "pdb", "pqa"},
		 "application/vnd.lotus-organizer": []string{"org"},
		 "application/vnd.yamaha.openscoreformat": []string{"osf"},
		 "application/vnd.yamaha.openscoreformat.osfpvg+xml": []string{"osfpvg"},
		 "application/vnd.oasis.opendocument.chart-template": []string{"otc"},
		 "application/x-font-otf": []string{"otf"},
		 "application/vnd.oasis.opendocument.graphics-template": []string{"otg"},
		 "application/vnd.oasis.opendocument.text-web": []string{"oth"},
		 "application/vnd.oasis.opendocument.image-template": []string{"oti"},
		 "application/vnd.oasis.opendocument.presentation-template": []string{"otp"},
		 "application/vnd.oasis.opendocument.spreadsheet-template": []string{"ots"},
		 "application/vnd.oasis.opendocument.text-template": []string{"ott"},
		 "application/oxps": []string{"oxps"},
		 "application/vnd.openofficeorg.extension": []string{"oxt"},
		 "application/pkcs10": []string{"p10"},
		 "application/x-pkcs12": []string{"p12", "pfx"},
		 "application/x-pkcs7-certificates": []string{"p7b", "spc"},
		 "application/pkcs7-mime": []string{"p7c", "p7m"},
		 "application/x-pkcs7-certreqresp": []string{"p7r"},
		 "application/pkcs7-signature": []string{"p7s"},
		 "application/pkcs8": []string{"p8"},
		 "text/x-pascal": []string{"pas", "p"},
		 "application/vnd.pawaafile": []string{"paw"},
		 "application/vnd.powerbuilder6": []string{"pbd"},
		 "image/x-portable-bitmap": []string{"pbm"},
		 "application/x-font-pcf": []string{"pcf"},
		 "application/vnd.hp-pcl": []string{"pcl"},
		 "application/vnd.hp-pclxl": []string{"pclxl"},
		 "image/x-pict": []string{"pct", "pic"},
		 "application/vnd.curl.pcurl": []string{"pcurl"},
		 "image/x-pcx": []string{"pcx"},
		 "application/pdf": []string{"pdf"},
		 "application/font-tdpfr": []string{"pfr"},
		 "image/x-portable-graymap": []string{"pgm"},
		 "application/x-chess-pgn": []string{"pgn"},
		 "application/pgp-encrypted": []string{"pgp"},
		 "application/pkixcmp": []string{"pki"},
		 "application/pkix-pkipath": []string{"pkipath"},
		 "application/vnd.3gpp.pic-bw-large": []string{"plb"},
		 "application/vnd.mobius.plc": []string{"plc"},
		 "application/vnd.pocketlearn": []string{"plf"},
		 "application/pls+xml": []string{"pls"},
		 "application/vnd.ctc-posml": []string{"pml"},
		 "image/png": []string{"png"},
		 "image/x-portable-anymap": []string{"pnm"},
		 "application/vnd.macports.portpkg": []string{"portpkg"},
		 "application/vnd.ms-powerpoint": []string{"ppt", "pps", "pot"},
		 "application/vnd.ms-powerpoint.template.macroenabled.12": []string{"potm"},
		 "application/vnd.openxmlformats-officedocument.presentationml.template": []string{"potx"},
		 "application/vnd.ms-powerpoint.addin.macroenabled.12": []string{"ppam"},
		 "application/vnd.cups-ppd": []string{"ppd"},
		 "image/x-portable-pixmap": []string{"ppm"},
		 "application/vnd.ms-powerpoint.slideshow.macroenabled.12": []string{"ppsm"},
		 "application/vnd.openxmlformats-officedocument.presentationml.slideshow": []string{"ppsx"},
		 "application/vnd.ms-powerpoint.presentation.macroenabled.12": []string{"pptm"},
		 "application/vnd.openxmlformats-officedocument.presentationml.presentation": []string{"pptx"},
		 "application/vnd.lotus-freelance": []string{"pre"},
		 "application/pics-rules": []string{"prf"},
		 "application/vnd.3gpp.pic-bw-small": []string{"psb"},
		 "image/vnd.adobe.photoshop": []string{"psd"},
		 "application/x-font-linux-psf": []string{"psf"},
		 "application/pskc+xml": []string{"pskcxml"},
		 "application/vnd.pvi.ptid1": []string{"ptid"},
		 "application/x-mspublisher": []string{"pub"},
		 "application/vnd.3gpp.pic-bw-var": []string{"pvb"},
		 "application/vnd.3m.post-it-notes": []string{"pwn"},
		 "audio/vnd.ms-playready.media.pya": []string{"pya"},
		 "video/vnd.ms-playready.media.pyv": []string{"pyv"},
		 "application/vnd.epson.quickanime": []string{"qam"},
		 "application/vnd.intu.qbo": []string{"qbo"},
		 "application/vnd.intu.qfx": []string{"qfx"},
		 "application/vnd.publishare-delta-tree": []string{"qps"},
		 "application/vnd.quark.quarkxpress": []string{"qwd", "qwt", "qxb", "qxd", "qxl", "qxt"},
		 "audio/x-pn-realaudio": []string{"ra", "ram"},
		 "application/x-rar-compressed": []string{"rar"},
		 "image/x-cmu-raster": []string{"ras"},
		 "application/vnd.ipunplugged.rcprofile": []string{"rcprofile"},
		 "application/rdf+xml": []string{"rdf"},
		 "application/vnd.data-vision.rdz": []string{"rdz"},
		 "application/vnd.businessobjects": []string{"rep"},
		 "application/x-dtbresource+xml": []string{"res"},
		 "image/x-rgb": []string{"rgb"},
		 "application/reginfo+xml": []string{"rif"},
		 "audio/vnd.rip": []string{"rip"},
		 "application/x-research-info-systems": []string{"ris"},
		 "application/resource-lists+xml": []string{"rl"},
		 "image/vnd.fujixerox.edmics-rlc": []string{"rlc"},
		 "application/resource-lists-diff+xml": []string{"rld"},
		 "application/vnd.rn-realmedia": []string{"rm"},
		 "audio/x-pn-realaudio-plugin": []string{"rmp"},
		 "application/vnd.jcp.javame.midlet-rms": []string{"rms"},
		 "application/vnd.rn-realmedia-vbr": []string{"rmvb"},
		 "application/relax-ng-compact-syntax": []string{"rnc"},
		 "application/rpki-roa": []string{"roa"},
		 "application/vnd.cloanto.rp9": []string{"rp9"},
		 "application/vnd.nokia.radio-presets": []string{"rpss"},
		 "application/vnd.nokia.radio-preset": []string{"rpst"},
		 "application/sparql-query": []string{"rq"},
		 "application/rls-services+xml": []string{"rs"},
		 "application/rsd+xml": []string{"rsd"},
		 "application/rss+xml": []string{"rss"},
		 "application/rtf": []string{"rtf"},
		 "text/richtext": []string{"rtx"},
		 "audio/s3m": []string{"s3m"},
		 "application/vnd.yamaha.smaf-audio": []string{"saf"},
		 "application/sbml+xml": []string{"sbml"},
		 "application/vnd.ibm.secure-container": []string{"sc"},
		 "application/x-msschedule": []string{"scd"},
		 "application/vnd.lotus-screencam": []string{"scm"},
		 "application/scvp-cv-request": []string{"scq"},
		 "application/scvp-cv-response": []string{"scs"},
		 "text/vnd.curl.scurl": []string{"scurl"},
		 "application/vnd.stardivision.draw": []string{"sda"},
		 "application/vnd.stardivision.calc": []string{"sdc"},
		 "application/vnd.stardivision.impress": []string{"sdd"},
		 "application/vnd.solent.sdkm+xml": []string{"sdkd", "sdkm"},
		 "application/sdp": []string{"sdp"},
		 "application/vnd.stardivision.writer": []string{"sdw", "sgl", "vor"},
		 "application/vnd.seemail": []string{"see"},
		 "application/vnd.sema": []string{"sema"},
		 "application/vnd.semd": []string{"semd"},
		 "application/vnd.semf": []string{"semf"},
		 "application/java-serialized-object": []string{"ser"},
		 "application/set-payment-initiation": []string{"setpay"},
		 "application/set-registration-initiation": []string{"setreg"},
		 "application/vnd.hydrostatix.sof-data": []string{"sfd-hdstx"},
		 "application/vnd.spotfire.sfs": []string{"sfs"},
		 "text/x-sfv": []string{"sfv"},
		 "image/sgi": []string{"sgi"},
		 "text/sgml": []string{"sgml", "sgm"},
		 "application/x-sh": []string{"sh"},
		 "application/x-shar": []string{"shar"},
		 "application/shf+xml": []string{"shf"},
		 "image/x-mrsid-image": []string{"sid"},
		 "audio/silk": []string{"sil"},
		 "application/vnd.symbian.install": []string{"sis", "sisx"},
		 "application/x-stuffit": []string{"sit"},
		 "application/x-stuffitx": []string{"sitx"},
		 "application/vnd.koan": []string{"skd", "skm", "skp", "skt"},
		 "application/vnd.ms-powerpoint.slide.macroenabled.12": []string{"sldm"},
		 "application/vnd.openxmlformats-officedocument.presentationml.slide": []string{"sldx"},
		 "application/vnd.epson.salt": []string{"slt"},
		 "application/vnd.stepmania.stepchart": []string{"sm"},
		 "application/vnd.stardivision.math": []string{"smf"},
		 "application/smil+xml": []string{"smi", "smil"},
		 "video/x-smv": []string{"smv"},
		 "application/vnd.stepmania.package": []string{"smzip"},
		 "application/x-font-snf": []string{"snf"},
		 "application/vnd.yamaha.smaf-phrase": []string{"spf"},
		 "application/x-futuresplash": []string{"spl"},
		 "text/vnd.in3d.spot": []string{"spot"},
		 "application/scvp-vp-response": []string{"spp"},
		 "application/scvp-vp-request": []string{"spq"},
		 "application/x-sql": []string{"sql"},
		 "application/x-wais-source": []string{"src"},
		 "application/x-subrip": []string{"srt"},
		 "application/sru+xml": []string{"sru"},
		 "application/sparql-results+xml": []string{"srx"},
		 "application/ssdl+xml": []string{"ssdl"},
		 "application/vnd.kodak-descriptor": []string{"sse"},
		 "application/vnd.epson.ssf": []string{"ssf"},
		 "application/ssml+xml": []string{"ssml"},
		 "application/vnd.sailingtracker.track": []string{"st"},
		 "application/vnd.sun.xml.calc.template": []string{"stc"},
		 "application/vnd.sun.xml.draw.template": []string{"std"},
		 "application/vnd.wt.stf": []string{"stf"},
		 "application/vnd.sun.xml.impress.template": []string{"sti"},
		 "application/hyperstudio": []string{"stk"},
		 "application/vnd.ms-pki.stl": []string{"stl"},
		 "application/vnd.pg.format": []string{"str"},
		 "application/vnd.sun.xml.writer.template": []string{"stw"},
		 "image/vnd.dvb.subtitle": []string{"sub"},
		 "text/vnd.dvb.subtitle": []string{"sub"},
		 "application/vnd.sus-calendar": []string{"sus", "susp"},
		 "application/x-sv4cpio": []string{"sv4cpio"},
		 "application/x-sv4crc": []string{"sv4crc"},
		 "application/vnd.dvb.service": []string{"svc"},
		 "application/vnd.svd": []string{"svd"},
		 "image/svg+xml": []string{"svg", "svgz"},
		 "application/x-shockwave-flash": []string{"swf"},
		 "application/vnd.aristanetworks.swi": []string{"swi"},
		 "application/vnd.sun.xml.calc": []string{"sxc"},
		 "application/vnd.sun.xml.draw": []string{"sxd"},
		 "application/vnd.sun.xml.writer.global": []string{"sxg"},
		 "application/vnd.sun.xml.impress": []string{"sxi"},
		 "application/vnd.sun.xml.math": []string{"sxm"},
		 "application/vnd.sun.xml.writer": []string{"sxw"},
		 "application/x-t3vm-image": []string{"t3"},
		 "application/vnd.mynfc": []string{"taglet"},
		 "application/vnd.tao.intent-module-archive": []string{"tao"},
		 "application/x-tar": []string{"tar"},
		 "application/vnd.3gpp2.tcap": []string{"tcap"},
		 "application/x-tcl": []string{"tcl"},
		 "application/vnd.smart.teacher": []string{"teacher"},
		 "application/tei+xml": []string{"tei", "teicorpus"},
		 "application/x-tex": []string{"tex"},
		 "application/x-texinfo": []string{"texi", "texinfo"},
		 "application/thraud+xml": []string{"tfi"},
		 "application/x-tex-tfm": []string{"tfm"},
		 "image/x-tga": []string{"tga"},
		 "application/vnd.ms-officetheme": []string{"thmx"},
		 "image/tiff": []string{"tiff", "tif"},
		 "application/vnd.tmobile-livetv": []string{"tmo"},
		 "application/x-bittorrent": []string{"torrent"},
		 "application/vnd.groove-tool-template": []string{"tpl"},
		 "application/vnd.trid.tpt": []string{"tpt"},
		 "application/vnd.trueapp": []string{"tra"},
		 "application/x-msterminal": []string{"trm"},
		 "application/timestamped-data": []string{"tsd"},
		 "text/tab-separated-values": []string{"tsv"},
		 "application/x-font-ttf": []string{"ttc", "ttf"},
		 "text/turtle": []string{"ttl"},
		 "application/vnd.simtech-mindmapper": []string{"twd", "twds"},
		 "application/vnd.genomatix.tuxedo": []string{"txd"},
		 "application/vnd.mobius.txf": []string{"txf"},
		 "application/vnd.ufdl": []string{"ufd", "ufdl"},
		 "application/x-glulx": []string{"ulx"},
		 "application/vnd.umajin": []string{"umj"},
		 "application/vnd.unity": []string{"unityweb"},
		 "application/vnd.uoml+xml": []string{"uoml"},
		 "text/uri-list": []string{"uris", "uri", "urls"},
		 "application/x-ustar": []string{"ustar"},
		 "application/vnd.uiq.theme": []string{"utz"},
		 "text/x-uuencode": []string{"uu"},
		 "audio/vnd.dece.audio": []string{"uva", "uvva"},
		 "application/vnd.dece.data": []string{"uvd", "uvf", "uvvd", "uvvf"},
		 "image/vnd.dece.graphic": []string{"uvg", "uvi", "uvvg", "uvvi"},
		 "video/vnd.dece.hd": []string{"uvh", "uvvh"},
		 "video/vnd.dece.mobile": []string{"uvm", "uvvm"},
		 "video/vnd.dece.pd": []string{"uvp", "uvvp"},
		 "video/vnd.dece.sd": []string{"uvs", "uvvs"},
		 "application/vnd.dece.ttml+xml": []string{"uvt", "uvvt"},
		 "video/vnd.uvvu.mp4": []string{"uvu"},
		 "video/vnd.dece.video": []string{"uvv", "uvvv"},
		 "application/vnd.dece.unspecified": []string{"uvx","uvvx"},
		 "application/vnd.dece.zip": []string{"uvz", "uvvz"},
		 "text/vcard": []string{"vcard"},
		 "application/x-cdlink": []string{"vcd"},
		 "text/x-vcard": []string{"vcf"},
		 "application/vnd.groove-vcard": []string{"vcg"},
		 "text/x-vcalendar": []string{"vcs"},
		 "application/vnd.vcx": []string{"vcx"},
		 "application/vnd.visionary": []string{"vis"},
		 "video/vnd.vivo": []string{"viv"},
		 "video/x-ms-vob": []string{"vob"},
		 "model/vrml": []string{"vrml", "wrl", "wrl"},
		 "application/vnd.visio": []string{"vsd", "vss", "vst", "vsw"},
		 "application/vnd.vsf": []string{"vsf"},
		 "model/vnd.vtu": []string{"vtu"},
		 "application/voicexml+xml": []string{"vxml"},
		 "application/x-doom": []string{"wad"},
		 "audio/x-wav": []string{"wav"},
		 "audio/wav": []string{"wav"},
		 "audio/x-ms-wax": []string{"wax"},
		 "image/vnd.wap.wbmp": []string{"wbmp"},
		 "application/vnd.criticaltools.wbs+xml": []string{"wbs"},
		 "application/vnd.wap.wbxml": []string{"wbxml"},
		 "application/vnd.ms-works": []string{"wcm", "wdb", "wks", "wps"},
		 "image/vnd.ms-photo": []string{"wdp"},
		 "audio/webm": []string{"weba"},
		 "video/webm": []string{"webm"},
		 "image/webp": []string{"webp"},
		 "application/vnd.pmi.widget": []string{"wg"},
		 "application/widget": []string{"wgt"},
		 "audio/x-ms-wma": []string{"wma"},
		 "application/x-ms-wmd": []string{"wmd"},
		 "application/vnd.wap.wmlc": []string{"wmlc"},
		 "application/vnd.wap.wmlscriptc": []string{"wmlsc"},
		 "text/vnd.wap.wmlscript": []string{"wmls"},
		 "text/vnd.wap.wml": []string{"wml"},
		 "video/x-ms-wm": []string{"wm"},
		 "video/x-ms-wmv": []string{"wmv"},
		 "video/x-ms-wmx": []string{"wmx"},
		 "application/x-ms-wmz": []string{"wmz"},
		 "application/font-woff": []string{"woff"},
		 "application/vnd.wordperfect": []string{"wpd"},
		 "application/vnd.ms-wpl": []string{"wpl"},
		 "application/vnd.wqd": []string{"wqd"},
		 "application/x-mswrite": []string{"wri"},
		 "application/wsdl+xml": []string{"wsdl"},
		 "application/wspolicy+xml": []string{"wspolicy"},
		 "application/vnd.webturbo": []string{"wtb"},
		 "video/x-ms-wvx": []string{"wvx"},
		 "model/x3d+binary": []string{"x3db", "x3dbz"},
		 "model/x3d+xml": []string{"x3d", "x3dz"},
		 "model/x3d+vrml": []string{"x3dv", "x3dvz"},
		 "application/xaml+xml": []string{"xaml"},
		 "application/x-silverlight-app": []string{"xap"},
		 "application/vnd.xara": []string{"xar"},
		 "application/x-ms-xbap": []string{"xbap"},
		 "application/vnd.fujixerox.docuworks.binder": []string{"xbd"},
		 "image/x-xbitmap": []string{"xbm"},
		 "application/xcap-diff+xml": []string{"xdf"},
		 "application/vnd.syncml.dm+xml": []string{"xdm"},
		 "application/vnd.adobe.xdp+xml": []string{"xdp"},
		 "application/dssc+xml": []string{"xdssc"},
		 "application/vnd.fujixerox.docuworks": []string{"xdw"},
		 "application/xenc+xml": []string{"xenc"},
		 "application/patch-ops-error+xml": []string{"xer"},
		 "application/vnd.adobe.xfdf": []string{"xfdf"},
		 "application/vnd.xfdl": []string{"xfdl"},
		 "application/xhtml+xml": []string{"xhtml", "xht"},
		 "image/vnd.xiff": []string{"xif"},
		 "application/vnd.ms-excel.addin.macroenabled.12": []string{"xlam"},
		 "application/vnd.ms-excel": []string{"xls", "xla", "xlc", "xlm", "xlt", "xlw"},
		 "application/x-xliff+xml": []string{"xlf"},
		 "application/vnd.ms-excel.sheet.binary.macroenabled.12": []string{"xlsb"},
		 "application/vnd.ms-excel.sheet.macroenabled.12": []string{"xlsm"},
		 "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": []string{"xlsx"},
		 "application/vnd.ms-excel.template.macroenabled.12": []string{"xltm"},
		 "application/vnd.openxmlformats-officedocument.spreadsheetml.template": []string{"xltx"},
		 "audio/xm": []string{"xm"},
		 "application/xml": []string{"xml", "xsl"},
		 "application/vnd.olpc-sugar": []string{"xo"},
		 "application/xop+xml": []string{"xop"},
		 "application/x-xpinstall": []string{"xpi"},
		 "application/xproc+xml": []string{"xpl"},
		 "image/x-xpixmap": []string{"xpm"},
		 "application/vnd.is-xpr": []string{"xpr"},
		 "application/vnd.ms-xpsdocument": []string{"xps"},
		 "application/vnd.intercon.formnet": []string{"xpw", "xpx"},
		 "application/xslt+xml": []string{"xslt"},
		 "application/vnd.syncml+xml": []string{"xsm"},
		 "application/xspf+xml": []string{"xspf"},
		 "application/vnd.mozilla.xul+xml": []string{"xul"},
		 "image/x-xwindowdump": []string{"xwd"},
		 "chemical/x-xyz": []string{"xyz"},
		 "application/x-xz": []string{"xz"},
		 "application/yang": []string{"yang"},
		 "application/yin+xml": []string{"yin"},
		 "application/x-zmachine": []string{"z1", "z2", "z3", "z4", "z5", "z6", "z7", "z8"},
		 "application/vnd.zzazz.deck+xml": []string{"zaz"},
		 "application/zip": []string{"zip"},
		 "application/vnd.zul": []string{"zir", "zirz"},
		 "application/vnd.handheld-entertainment+xml": []string{"zmm"},
	}
	return mimeToExt
}