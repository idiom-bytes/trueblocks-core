/*-------------------------------------------------------------------------------------------
 * qblocks - fast, easily-accessible, fully-decentralized data from blockchains
 * copyright (c) 2016, 2021 TrueBlocks, LLC (http://trueblocks.io)
 *
 * This program is free software: you may redistribute it and/or modify it under the terms
 * of the GNU General Public License as published by the Free Software Foundation, either
 * version 3 of the License, or (at your option) any later version. This program is
 * distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
 * the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details. You should have received a copy of the GNU General
 * Public License along with this program. If not, see http://www.gnu.org/licenses/.
 *-------------------------------------------------------------------------------------------*/
#include "etherlib.h"

int firstCout = true;
CAbi abi_spec;

//----------------------------------------------------------------
bool visitTransaction(CTransaction& trans, void* data) {
    abi_spec.loadAbiFromEtherscan(trans.to);
    for (auto log : trans.receipt.logs)
        abi_spec.loadAbiFromEtherscan(log.address);
    abi_spec.articulateTransaction(&trans);

    if(!firstCout)
        cout << ",";
    cout << "  ";
    indent();
    trans.toJson(cout);
    unindent();

    firstCout = false;
    return true;
}

//----------------------------------------------------------------
bool visitBlock(CBlock& block, void* data) {
    return block.forEveryTransaction(visitTransaction, data);
}

// Options Main ----------------------------------------------------------------
int main(int argc, const char* argv[]) {
    loadEnvironmentPaths();
    etherlib_init(quickQuitHandler);

    int start = atoi(argv[1]);
    int nBlocks = atoi(argv[2]);

    manageFields(defHide, false);
    manageFields(defShow, true);
    manageFields("CParameter:strDefault", false);  // hide
    manageFields("CTransaction:price", false);     // hide
    manageFields("CFunction:outputs", true);                                               // show
    manageFields("CTransaction:input", true);                                                  // show
    manageFields("CLogEntry:data,topics", true);                                               // show
    manageFields("CTrace: blockHash, blockNumber, transactionHash, transactionIndex", false);  // hide
    abi_spec.loadAbisFromKnown();

    // Core logic
    uint32_t counter = 0;
    cout << "[";
    forEveryBlock(visitBlock, &counter, start, nBlocks+1, 1);
    cout << "]";

    etherlib_cleanup();

    return 0;

}